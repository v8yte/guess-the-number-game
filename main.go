package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main(){
    // 乱数生成期の初期化
    rand.Seed(time.Now().UnixNano())

    // 最小数（n）と最大数（m）
    var n,m,r,u int
    fmt.Print("最小値:")
    if _,err := fmt.Scanf("%d",&n); err != nil {
        fmt.Println("無効な入力です。")
        return
    }
    fmt.Print("最大値:")
    if _,err := fmt.Scanf("%d",&m); err != nil {
        fmt.Println("無効な入力です。")
        return
    }
    
    if n > m {
        fmt.Println("最小値は最大値以下にしてください。")
        return
    }

    // nからmの範囲の乱数を生成
    if n == m {
        r = n
    }else{
        r = rand.Intn(m-n) + n
    }

    // ユーザーの予測値との比較
    for i := 0; i < 3; i++ {

        fmt.Print("あなたの予想:")
        if _,err := fmt.Scanf("%d",&u); err != nil {
            fmt.Println("無効な入力です。")
            return
        }

        if u == r {
            fmt.Println("正解！")
            return 
        }
        fmt.Println("ハズレ")
    }
    fmt.Printf("答えは%dでした。\n",r)
}
