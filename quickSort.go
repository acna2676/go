package main

import (
    "fmt"
    "time"
    "math/rand"
)

func sort(values []int) (ret []int) {
    // 要素数が 1 以下の配列はそれ以上細分化してソートする必要がない!
    if len(values) < 2 {
        return values
    }

    // 配列の先頭をピボット (基準値) に選ぶ
    pivot := values[0]

    // ピボットを基準にして値の大小で配列をふたつに分割する
    left := []int{}
    right := []int{}
    for _, v := range values[1:] {
        if v > pivot {
            right = append(right, v)
        } else {
            left = append(left, v)
        }
    }

    // 分割した配列をそれぞれ再帰的にソートする
    left = sort(left)
    right = sort(right)

    /*
    left(小さい) + pivot(基準値) + right(大きい) の順番で配列を組み立てる
    ここが実際のソート処理
    */
    ret = append(left, pivot)
    ret = append(ret, right...)

    return
}

func main() {
    // UNIX 時間をシードにして乱数生成器を用意する
    t := time.Now().Unix()
    s := rand.NewSource(t)
    r := rand.New(s)

    // ランダムな値の入った配列を作る
    N := 10
    values := make([]int, N)
    for i := 0; i < N; i++ {
        values[i] = r.Intn(N)
    }

    // ソートして結果を出力する
    sortedValues := sort(values)
    fmt.Println(sortedValues)
}
