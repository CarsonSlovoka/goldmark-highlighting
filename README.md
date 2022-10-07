<p align="center">
  <a href="http://golang.org">
      <img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg" alt="Made with Go">
  </a>

  <img src="https://img.shields.io/github/go-mod/go-version/CarsonSlovoka/goldmark-highlighting?filename=v2%2Fgo.mod" alt="Go Version">

  <a href="https://GitHub.com/CarsonSlovoka/goldmark-highlighting/releases/">
      <img src="https://img.shields.io/github/release/CarsonSlovoka/goldmark-highlighting" alt="Latest release">
  </a>
  <a href="https://github.com/CarsonSlovoka/goldmark-highlighting/blob/master/LICENSE">
      <img src="https://img.shields.io/github/license/CarsonSlovoka/goldmark-highlighting.svg" alt="License">
  </a>
</p>

# goldmark-highlighting

這是一個[goldmark](https://github.com/yuin/goldmark)的擴充功能

## Features

- 支持屬性調整:

  > \`\`\`go {style="vim" hls=[1, 3-5] base=3 linenos=table tw=2}

  名稱使用精簡的方式，縮減md檔案的大小，各個項目的內容可以參考[Attributes](https://github.com/CarsonSlovoka/goldmark-highlighting/blob/2451b2d1fe43790cb44537bff1bdb27c27e48e4d/v2/renderer-method.go#L20-L28)

  或者參考[attr的測試](https://github.com/CarsonSlovoka/goldmark-highlighting/tree/2451b2d/v2/testData/attr)，以了解用法

## Tutorial

如果您只想要學習goldmark該怎麼寫code-block的擴展，了解其基本精神，可以參考以下2個歷史紀錄，應該就能掌握
1. [init Extender](https://github.com/CarsonSlovoka/goldmark-highlighting/commit/10479a0204a2ad01b9aec8ab19fcc4d7c1b6a5c2)
2. [可運行首版](https://github.com/CarsonSlovoka/goldmark-highlighting/commit/fa41260c25144ee39760b67e3eccbe0694eb0975)
   - [測試](https://github.com/CarsonSlovoka/goldmark-highlighting/commit/89d8ef7a4967f2f505a2e86f7e88620439ea371a)

## 參考資料

- [github.com/yuin/goldmark-highlighting](https://github.com/yuin/goldmark-highlighting)
- [github.com/alecthomas/chroma](https://github.com/alecthomas/chroma) 掌管主題顏色的細節，例如{背景、數字、註解、字串...}[等相關屬性](https://github.com/alecthomas/chroma/blob/6138519d55582350e5dec0147cb8f5ddcb78f8cf/styles/swapoff.go#L8-L25)
  - Style
    - [Chroma Style Gallery總覽](https://xyproto.github.io/splash/docs/all.html)
    - [Chroma Style Gallery](https://xyproto.github.io/splash/docs/)
    - [Source](https://github.com/alecthomas/chroma/tree/3f86ac7/styles): 這邊每一個檔案都表示一種style
