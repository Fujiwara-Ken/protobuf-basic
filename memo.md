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
