# 六十進制

六十進制，可以以天干地支表示。
也可以以六十四卦的前六十卦表示。

只是一些想法，目前還不具有實際意義。

好處：

* 兼容十進制，特別是處理小數時，不會丟失精度
* 兼容二進制
* 方便計算一些常見分數，如1/2, 1/3, 1/4, 1/5, 1/6, 1/10，1/20, 1/30, 1/60
* 方便時間和角度的計算
* 表述一些特定文化，如天干地支、六十四卦等
* 更加節省存儲空間
* 可以規定一位為一個字節，消除“位”、“字節”可能引起的歧義，以後只以字節表示

目前存在的問題：

* 目前的計算機系統還是以二進制為基礎，包括存儲、計算、傳輸等。六十進制的優點可能無法表現出來。

如何兼容：

* 位位對應，也就是一個六十進制位對應一個二進制位，這種方式最簡單，但浪費較大。
* 六十四位二進制，對應十一位六十進制。六十四位二進制，可以包含目前常用的INT64、DOUBLE、INT32、FLOAT等，單個字符（絕大多數文字）
* 每個字符最多4個字節（類似uft-8）
* 二進制模擬六十進制，由八位二進制表示
* 字符串採用[UTF-60](#utf-60)
* 數字。保持以二進制方式進行計算

## 天干地支

| 10    | 20    | 30    | 40    | 50    | 60    |
|-------|-------|-------|-------|-------|-------|
| 01 甲子 | 11 甲戌 | 21 甲申 | 31 甲午 | 41 甲辰 | 51 甲寅 |
| 02 乙丑 | 12 乙亥 | 22 乙酉 | 32 乙未 | 42 乙巳 | 52 乙卯 |
| 03 丙寅 | 13 丙子 | 23 丙戌 | 33 丙申 | 43 丙午 | 53 丙辰 |
| 04 丁卯 | 14 丁丑 | 24 丁亥 | 34 丁酉 | 44 丁未 | 54 丁巳 |
| 05 戊辰 | 15 戊寅 | 25 戊子 | 35 戊戌 | 45 戊申 | 55 戊午 |
| 06 己巳 | 16 己卯 | 26 己丑 | 36 己亥 | 46 己酉 | 56 己未 |
| 07 庚午 | 17 庚辰 | 27 庚寅 | 37 庚子 | 47 庚戌 | 57 庚申 |
| 08 辛未 | 18 辛巳 | 28 辛卯 | 38 辛丑 | 48 辛亥 | 58 辛酉 |
| 09 壬申 | 19 壬午 | 29 壬辰 | 39 壬寅 | 49 壬子 | 59 壬戌 |
| 10 癸酉 | 20 癸未 | 30 癸巳 | 40 癸卯 | 50 癸丑 | 60 癸亥 |

## 六十四卦

| 10   | 20   | 30   | 40   | 50   | 60   |
|------|------|------|------|------|------|
| 01 ䷁ | 11 ䷦ | 21 ䷧ | 31 ䷛ | 41 ䷣ | 51 ䷻ |
| 02 ䷖ | 12 ䷴ | 22 ䷿ | 32 ䷫ | 42 ䷕ | 52 ䷼ |
| 03 ䷇ | 13 ䷽ | 23 ䷮ | 33 ䷗ | 43 ䷾ | 53 ䷵ |
| 04 ䷓ | 14 ䷷ | 24 ䷅ | 34 ䷚ | 44 ䷤ | 54 ䷥ |
| 05 ䷏ | 15 ䷞ | 25 ䷭ | 35 ䷂ | 45 ䷶ | 55 ䷹ |
| 06 ䷢ | 16 ䷠ | 26 ䷑ | 36 ䷩ | 46 ䷝ | 56 ䷉ |
| 07 ䷬ | 17 ䷆ | 27 ䷯ | 37 ䷲ | 47 ䷰ | 57 ䷊ |
| 08 ䷋ | 18 ䷃ | 28 ䷸ | 38 ䷔ | 48 ䷌ | 58 ䷙ |
| 09 ䷎ | 19 ䷜ | 29 ䷟ | 39 ䷐ | 49 ䷒ | 59 ䷄ |
| 10 ䷳ | 20 ䷺ | 30 ䷱ | 40 ䷘ | 50 ䷨ | 60 ䷈ |

## UTF-60

基於UTF-8

一字節表

現在的UNICODE編碼，兼容了ASCII，其實並不是太好。
1. 部分字符已經不再常用，應該廢棄
2. 現代鍵盤的一些按鍵應該加入，很多地方都有用到。
3. 英文雖說很重要，但考慮到國際化，應該放到後面，比如和其他字母語言放到兩字節表

| 10                 | 20           | 30   | 40   | 50    | 60   |
|--------------------|--------------|------|------|-------|------|
| 01 Null            | 11 Shift     | 21 # | 31 - | 41 [  | 51 0 |
| 02 Tab             | 12 Up        | 22 $ | 32 . | 42 \  | 52 1 |
| 03 CAPS            | 13 Dwon      | 23 % | 33 / | 43 ]  | 53 2 |
| 04 New Line        | 14 Left      | 24 & | 34 : | 44 ^  | 54 3 |
| 05 New Page        | 15 Right     | 25 ' | 35 ; | 45 _  | 55 4 |
| 06 Carriage Return | 16 Escape    | 26 ( | 36 < | 46 `  | 56 5 |
| 07 Fn              | 17 Backspace | 27 ) | 37 = | 47 {  | 57 6 |
| 08 Control         | 18 Space     | 28 * | 38 > | 48 \| | 58 7 |
| 09 Alt             | 19 !         | 29 + | 39 ? | 49 }  | 59 8 |
| 10 Command         | 20 "         | 30 , | 40 @ | 50 ~  | 60 9 |

兩字節表

三字節表

四字節表