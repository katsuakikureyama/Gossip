# Gossip

WordPress の原本と履歴を受け継ぎ、Go による実行環境と Go 製プラグインを探る、実験的な WordPress フォークです。

Gossip is an independent, experimental WordPress fork exploring a Go-powered runtime and native Go plugins. It preserves the original WordPress source and history and is not affiliated with or endorsed by the WordPress project.

## 開発状況

**現在は研究・開発段階です。WordPress 全体を Go で動かせる状態ではありません。**

- `master` は WordPress の原本と、この紹介用 README を保持しています。
- Go 実装と検証コードは開発ブランチ・Draft PR で管理しています。この README の追加では、実装の PR をマージしていません。
- 開発ブランチでは、選定した WordPress 原文の関数を GoHP で Go に変換し、PHP と動作を比較しています。[13関数・33ケースの統合検証](https://github.com/katsuakikureyama/Gossip/pull/15)は限定した入力での結果であり、WordPress 全体の互換性を示すものではありません。
- WordPress コアの Go での起動、データベース連携、既存テーマ・プラグインの互換性、本番運用は未確認です。

## 関連プロジェクト・資料

- [Gossip の Pull Requests](https://github.com/katsuakikureyama/Gossip/pulls) — 開発中の変更と進捗
- [開発記録・検証の範囲](https://github.com/katsuakikureyama/Gossip/blob/da03747ec39c94a3ff07201dda4143184cdd7a57/docs/current-state-fork-baseline.md) — 13関数の統合検証時点の記録
- [開発用 README](https://github.com/katsuakikureyama/Gossip/blob/da03747ec39c94a3ff07201dda4143184cdd7a57/README.md) — 実装ブランチでのローカル検証手順
- [GoHP](https://github.com/katsuakikureyama/GoHP) — PHP の限定した構文・動作を Go に変換する別プロジェクト
- [Gossip-prototype](https://github.com/katsuakikureyama/Gossip-prototype) — 以前の独立 CMS プロトタイプ。現在の WordPress フォークとは別の実験です

## WordPress の原本について

`wp-admin/`、`wp-includes/`、`wp-content/` などの WordPress 原本は保持しています。
フォークの基点は WordPress コミット `54106f6c545fc8b881f6b39b7963002bece6c28c`（`7.2-beta1-64693`）です。開発スナップショットであり、安定した本番用リリースとして案内するものではありません。

WordPress 由来の [readme.html](readme.html) と [license.txt](license.txt) は変更せずに残しています。元のコードに付随するライセンス・著作権表示を保持し、この紹介文によって変更するものではありません。
