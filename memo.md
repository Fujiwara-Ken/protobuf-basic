# protocol buffers

## message

- 複数のフィールドを持つことができる型定義
  - それぞれのフィールドはスカラ型もしくはコンポジット型
- 各言語のコードとしてコンパイルした場合、構造体やクラスとして変換される
- 1 つの proto ファイルに複数の message 型を定義することも可能

= 1 などの =の後は代入ではなくタグ番号となる

```proto
message Employee {
  int32 id = 1;
  string name = 2;
  string email = 3;
}
```

スカラ型などの公式リファレンス
[https://developers.google.com/protocol-buffers/docs/overview]

## Tag

- Protocol Buffers ではフィールドはフィールド名ではなく、タグ番号によって識別される
- 重複は許されず、一意である必要がある
- タグの最小値は 1、最大値は 2^29 -1(536870911)
- 19000 ~ 19999 は Profocol Buffers の予約番号のため使用不可
- 1 ~ 15 番までは 1byte で表すことができるので、よく使うフィールドには 1~ 15 を割り当てるとパフォーマンスが向上する
- タグは連番にする必要はないので、あまり使わないフィールドはあえて 16 番以降を割り当てることも可能
- タグ番号を予約するなど、安全に Protocol Buffers を使用する方法も用意されている
