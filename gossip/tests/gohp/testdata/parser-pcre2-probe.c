/*
 * Fixed-profile PCRE2 probe for the WordPress block-parser token pattern.
 *
 * This is deliberately not a general-purpose regexp runner.  The Go test
 * supplies the extracted pattern and one raw subject per invocation, and
 * records the complete JSON response as comparison evidence.
 */
#define PCRE2_CODE_UNIT_WIDTH 8
#include <pcre2.h>

#include <ctype.h>
#include <errno.h>
#include <inttypes.h>
#include <limits.h>
#include <locale.h>
#include <stdbool.h>
#include <stdarg.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/stat.h>

#define INPUT_LIMIT 4096U
#define FIXED_CAPTURE_COUNT 5U
#define FIXED_NAME_COUNT 5U
#define FIXED_DEPTH_LIMIT 10000U
#define FIXED_HEAP_LIMIT_KIB (128U * 1024U)

struct bytes {
	uint8_t *data;
	size_t len;
};

struct builder {
	char *data;
	size_t len;
	size_t cap;
};

static void fail(const char *format, ...) {
	va_list args;
	fputs("parser-pcre2-probe: ", stderr);
	va_start(args, format);
	vfprintf(stderr, format, args);
	va_end(args);
	fputc('\n', stderr);
	exit(EXIT_FAILURE);
}

static void check_pcre2(int rc, const char *operation) {
	if (rc < 0) {
		fail("%s failed: %d", operation, rc);
	}
}

static void free_bytes(struct bytes *value) {
	free(value->data);
	value->data = NULL;
	value->len = 0;
}

static struct bytes read_input(const char *path, const char *label) {
	struct stat status;
	FILE *file;
	struct bytes value = {0};
	size_t nread;

	if (stat(path, &status) != 0) {
		fail("cannot stat %s %s: %s", label, path, strerror(errno));
	}
	if (!S_ISREG(status.st_mode)) {
		fail("%s must be a regular file: %s", label, path);
	}
	if (status.st_size < 0 || (uintmax_t)status.st_size > INPUT_LIMIT) {
		fail("%s exceeds %u bytes: %s", label, INPUT_LIMIT, path);
	}
	value.len = (size_t)status.st_size;
	value.data = malloc(value.len == 0 ? 1 : value.len);
	if (value.data == NULL) {
		fail("cannot allocate %s input", label);
	}
	file = fopen(path, "rb");
	if (file == NULL) {
		free_bytes(&value);
		fail("cannot open %s %s: %s", label, path, strerror(errno));
	}
	nread = value.len == 0 ? 0 : fread(value.data, 1, value.len, file);
	if (nread != value.len || ferror(file)) {
		fclose(file);
		free_bytes(&value);
		fail("cannot read %s %s", label, path);
	}
	if (fgetc(file) != EOF) {
		fclose(file);
		free_bytes(&value);
		fail("%s changed or exceeds %u bytes: %s", label, INPUT_LIMIT, path);
	}
	if (fclose(file) != 0) {
		free_bytes(&value);
		fail("cannot close %s %s", label, path);
	}
	return value;
}

static size_t parse_nonnegative(const char *text, const char *label) {
	char *end = NULL;
	uintmax_t value;
	size_t i;
	if (text[0] == '\0') {
		fail("%s must be a non-negative decimal integer", label);
	}
	for (i = 0; text[i] != '\0'; i++) {
		if (text[i] < '0' || text[i] > '9') {
			fail("%s must be a non-negative decimal integer", label);
		}
	}
	errno = 0;
	value = strtoumax(text, &end, 10);
	if (errno != 0 || end == text || *end != '\0' || value > SIZE_MAX) {
		fail("invalid %s: %s", label, text);
	}
	return (size_t)value;
}

static uint32_t parse_match_limit(const char *text) {
	size_t value = parse_nonnegative(text, "match limit");
	if (value != 1U && value != 100000U) {
		fail("match limit must be 1 or 100000");
	}
	return (uint32_t)value;
}

static void builder_init(struct builder *out) {
	out->cap = 1024;
	out->len = 0;
	out->data = malloc(out->cap);
	if (out->data == NULL) {
		fail("cannot allocate JSON output");
	}
	out->data[0] = '\0';
}

static void builder_free(struct builder *out) {
	free(out->data);
	out->data = NULL;
	out->len = 0;
	out->cap = 0;
}

static void builder_reserve(struct builder *out, size_t extra) {
	size_t required;
	char *grown;
	if (extra > SIZE_MAX - out->len - 1) {
		fail("JSON output too large");
	}
	required = out->len + extra + 1;
	if (required <= out->cap) {
		return;
	}
	while (out->cap < required) {
		if (out->cap > SIZE_MAX / 2) {
			out->cap = required;
			break;
		}
		out->cap *= 2;
	}
	grown = realloc(out->data, out->cap);
	if (grown == NULL) {
		fail("cannot grow JSON output");
	}
	out->data = grown;
}

static void builder_append(struct builder *out, const char *text) {
	size_t len = strlen(text);
	builder_reserve(out, len);
	memcpy(out->data + out->len, text, len + 1);
	out->len += len;
}

static void builder_appendf(struct builder *out, const char *format, ...) {
	va_list args;
	va_list copied;
	int needed;
	va_start(args, format);
	va_copy(copied, args);
	needed = vsnprintf(NULL, 0, format, copied);
	va_end(copied);
	if (needed < 0) {
		va_end(args);
		fail("cannot format JSON output");
	}
	builder_reserve(out, (size_t)needed);
	if (vsnprintf(out->data + out->len, out->cap - out->len, format, args) != needed) {
		va_end(args);
		fail("cannot format JSON output");
	}
	va_end(args);
	out->len += (size_t)needed;
}

static void builder_append_json_string(struct builder *out, const uint8_t *text, size_t len) {
	size_t i;
	builder_append(out, "\"");
	for (i = 0; i < len; i++) {
		uint8_t c = text[i];
		if (c == '\"' || c == '\\') {
			builder_appendf(out, "\\%c", c);
		} else if (c >= 0x20U && c <= 0x7eU) {
			builder_reserve(out, 1);
			out->data[out->len++] = (char)c;
			out->data[out->len] = '\0';
		} else {
			builder_appendf(out, "\\u%04x", c);
		}
	}
	builder_append(out, "\"");
}

static void builder_append_hex(struct builder *out, const uint8_t *data, size_t len) {
	static const char digits[] = "0123456789abcdef";
	size_t i;
	builder_append(out, "\"");
	builder_reserve(out, len * 2);
	for (i = 0; i < len; i++) {
		out->data[out->len++] = digits[data[i] >> 4];
		out->data[out->len++] = digits[data[i] & 0x0fU];
	}
	out->data[out->len] = '\0';
	builder_append(out, "\"");
}

static bool valid_pcre_name(const PCRE2_UCHAR *name, size_t len) {
	size_t i;
	if (len == 0) {
		return false;
	}
	for (i = 0; i < len; i++) {
		if (!(name[i] == '_' || (name[i] >= 'a' && name[i] <= 'z') ||
		      (name[i] >= 'A' && name[i] <= 'Z') || (name[i] >= '0' && name[i] <= '9'))) {
			return false;
		}
	}
	return true;
}

int main(int argc, char **argv) {
	struct bytes pattern = {0};
	struct bytes subject = {0};
	struct builder json = {0};
	pcre2_compile_context *compile_context = NULL;
	pcre2_match_context *match_context = NULL;
	pcre2_code *code = NULL;
	pcre2_match_data *match_data = NULL;
	PCRE2_UCHAR runtime_version[64] = {0};
	PCRE2_SPTR name_table = NULL;
	PCRE2_SIZE *ovector;
	uint32_t match_limit;
	uint32_t capture_count = 0;
	uint32_t name_count = 0;
	uint32_t name_entry_size = 0;
	uint32_t effective_options = 0;
	uint32_t effective_newline = 0;
	uint32_t effective_bsr = 0;
	uint32_t ovector_count;
	int error_code;
	PCRE2_SIZE error_offset;
	int return_code;
	size_t offset;
	size_t i;

	if (argc != 5) {
		fail("usage: %s PATTERN-BODY-FILE SUBJECT-FILE OFFSET MATCH-LIMIT", argv[0]);
	}
	if (setlocale(LC_ALL, "C") == NULL) {
		fail("cannot set LC_ALL=C");
	}
	if (PCRE2_MAJOR != 10 || PCRE2_MINOR != 47) {
		fail("PCRE2 headers must be 10.47, found %d.%d", PCRE2_MAJOR, PCRE2_MINOR);
	}
	check_pcre2(pcre2_config(PCRE2_CONFIG_VERSION, runtime_version), "pcre2_config(version)");
	if (strncmp((const char *)runtime_version, "10.47", 5) != 0 ||
	    (runtime_version[5] != '\0' && runtime_version[5] != ' ')) {
		fail("PCRE2 runtime must be 10.47, found %s", (const char *)runtime_version);
	}

	pattern = read_input(argv[1], "pattern");
	subject = read_input(argv[2], "subject");
	offset = parse_nonnegative(argv[3], "offset");
	match_limit = parse_match_limit(argv[4]);

	compile_context = pcre2_compile_context_create(NULL);
	if (compile_context == NULL) {
		fail("cannot create compile context");
	}
	check_pcre2(pcre2_set_newline(compile_context, PCRE2_NEWLINE_LF), "pcre2_set_newline");
	check_pcre2(pcre2_set_bsr(compile_context, PCRE2_BSR_UNICODE), "pcre2_set_bsr");
	code = pcre2_compile(pattern.data, pattern.len, PCRE2_DOTALL, &error_code, &error_offset, compile_context);
	if (code == NULL) {
		PCRE2_UCHAR message[256] = {0};
		(void)pcre2_get_error_message(error_code, message, sizeof(message));
		fail("pcre2_compile failed at %" PRIuMAX ": %d (%s)", (uintmax_t)error_offset,
		     error_code, (const char *)message);
	}
	check_pcre2(pcre2_pattern_info(code, PCRE2_INFO_CAPTURECOUNT, &capture_count),
	            "pcre2_pattern_info(capture count)");
	check_pcre2(pcre2_pattern_info(code, PCRE2_INFO_NAMECOUNT, &name_count),
	            "pcre2_pattern_info(name count)");
	check_pcre2(pcre2_pattern_info(code, PCRE2_INFO_NAMEENTRYSIZE, &name_entry_size),
	            "pcre2_pattern_info(name entry size)");
	check_pcre2(pcre2_pattern_info(code, PCRE2_INFO_NAMETABLE, &name_table),
	            "pcre2_pattern_info(name table)");
	check_pcre2(pcre2_pattern_info(code, PCRE2_INFO_ALLOPTIONS, &effective_options),
	            "pcre2_pattern_info(options)");
	check_pcre2(pcre2_pattern_info(code, PCRE2_INFO_NEWLINE, &effective_newline),
	            "pcre2_pattern_info(newline)");
	check_pcre2(pcre2_pattern_info(code, PCRE2_INFO_BSR, &effective_bsr),
	            "pcre2_pattern_info(bsr)");
	if (capture_count != FIXED_CAPTURE_COUNT || name_count != FIXED_NAME_COUNT || name_table == NULL ||
	    name_entry_size < 3) {
		fail("fixed pattern must have %u captures and %u names (got %u and %u)",
		     FIXED_CAPTURE_COUNT, FIXED_NAME_COUNT, capture_count, name_count);
	}
	if (effective_options != PCRE2_DOTALL || effective_newline != PCRE2_NEWLINE_LF ||
	    effective_bsr != PCRE2_BSR_UNICODE) {
		fail("compiled pattern profile differs from DOTALL/LF/BSR_UNICODE");
	}
	for (i = 0; i < name_count; i++) {
		PCRE2_SPTR entry = name_table + i * name_entry_size;
		size_t name_len = 0;
		while (name_len + 2 < name_entry_size && entry[2 + name_len] != '\0') {
			name_len++;
		}
		if (name_len + 2 >= name_entry_size || !valid_pcre_name(entry + 2, name_len)) {
			fail("invalid PCRE2 capture name in name table");
		}
	}

	match_context = pcre2_match_context_create(NULL);
	if (match_context == NULL) {
		fail("cannot create match context");
	}
	check_pcre2(pcre2_set_match_limit(match_context, match_limit), "pcre2_set_match_limit");
	check_pcre2(pcre2_set_depth_limit(match_context, FIXED_DEPTH_LIMIT), "pcre2_set_depth_limit");
	check_pcre2(pcre2_set_heap_limit(match_context, FIXED_HEAP_LIMIT_KIB), "pcre2_set_heap_limit");
	match_data = pcre2_match_data_create_from_pattern(code, NULL);
	if (match_data == NULL) {
		fail("cannot create match data");
	}
	ovector_count = pcre2_get_ovector_count(match_data);
	if (ovector_count != capture_count + 1U) {
		fail("match data capture count mismatch");
	}
	return_code = pcre2_match(code, subject.data, subject.len, offset, 0, match_data, match_context);
	if (return_code > 0 && (uint32_t)return_code > ovector_count) {
		fail("PCRE2 returned more captures than allocated match data");
	}
	ovector = pcre2_get_ovector_pointer(match_data);
	if (ovector == NULL) {
		fail("cannot obtain match ovector");
	}

	builder_init(&json);
	builder_append(&json, "{\"version\":");
	builder_append_json_string(&json, runtime_version, strlen((const char *)runtime_version));
	builder_appendf(&json,
	                ",\"header_version\":\"%d.%d\",\"jit\":false,\"match_limit\":%" PRIu32
	                ",\"depth_limit\":%u,\"heap_limit_kib\":%u,\"compile_options\":%" PRIu32
	                ",\"newline\":%" PRIu32 ",\"bsr\":%" PRIu32 ",\"return_code\":%d"
	                ",\"capture_count\":%" PRIu32 ",\"names\":[",
	                PCRE2_MAJOR, PCRE2_MINOR, match_limit, FIXED_DEPTH_LIMIT, FIXED_HEAP_LIMIT_KIB,
	                effective_options, effective_newline, effective_bsr, return_code, capture_count);
	for (i = 0; i < name_count; i++) {
		PCRE2_SPTR entry = name_table + i * name_entry_size;
		size_t name_len = 0;
		uint32_t number = ((uint32_t)entry[0] << 8) | entry[1];
		while (entry[2 + name_len] != '\0') {
			name_len++;
		}
		if (i != 0) {
			builder_append(&json, ",");
		}
		builder_appendf(&json, "{\"number\":%" PRIu32 ",\"name\":", number);
		builder_append_json_string(&json, entry + 2, name_len);
		builder_append(&json, "}");
	}
	builder_append(&json, "],\"captures\":[");
	if (return_code > 0) {
		for (i = 0; i < (size_t)return_code; i++) {
			PCRE2_SIZE start = ovector[2 * i];
			PCRE2_SIZE end = ovector[2 * i + 1];
			if (i != 0) {
				builder_append(&json, ",");
			}
			if (start == PCRE2_UNSET || end == PCRE2_UNSET) {
				builder_appendf(&json, "{\"number\":%zu,\"start\":-1,\"end\":-1,\"hex\":\"\"}", i);
				continue;
			}
			if (start > end || end > subject.len) {
				fail("PCRE2 returned invalid capture bounds");
			}
			builder_appendf(&json, "{\"number\":%zu,\"start\":%zu,\"end\":%zu,\"hex\":", i,
			                (size_t)start, (size_t)end);
			builder_append_hex(&json, subject.data + start, (size_t)(end - start));
			builder_append(&json, "}");
		}
	}
	builder_append(&json, "]}\n");
	if (fputs(json.data, stdout) == EOF || fflush(stdout) != 0) {
		fail("cannot write JSON output");
	}

	builder_free(&json);
	pcre2_match_data_free(match_data);
	pcre2_match_context_free(match_context);
	pcre2_code_free(code);
	pcre2_compile_context_free(compile_context);
	free_bytes(&subject);
	free_bytes(&pattern);
	return EXIT_SUCCESS;
}
