package main
import "fmt"
type node struct{
    val float64
    char byte
    next map[*node]bool
    inDim int
    outDim int
}

func hasCircle(c byte,holeMap *map[byte]*node,isOk *map[byte]bool,path *map[byte]bool,r *bool)bool {
    if !(*r){
        return false
    }

    if p,ok := (*isOk)[c];ok{
        if p{
            return true
        }
        (*r) = false
    }

    if (*path)[c]{
        (*isOk)[c] = false
        (*r) = false
        return false
    }

    (*path)[c] = true
    ptr := (*holeMap)[c]
    for s,_ := range ptr.next{
        if !hasCircle(s.char,holeMap,isOk,path,r){
            (*isOk)[c] = false
            (*r) = false
            return false
        }
    }

    (*path)[c] =false
    return true
}

func diff(s0,s1 string)(byte,byte){
    l0 := len(s0)
    l1 := len(s1)

    l := l0
    if l > l1{
        l = l1
    }
    //fmt.Println("-----")
    //fmt.Println(s0,s1)
    for i:=0;i<l;i++{
        //fmt.Println(string(s0[i]) + " "+string(s1[i]))
        if s0[i] == s1[i]{
            continue
        }
        return s0[i],s1[i]
    }
    return 1,1
}
func alienOrder(words []string) string {
    holeMap := make(map[byte]*node)
    
    // 建图的节点
    for _,word := range words{
        for v := range word{
            c := word[v]
            if _,ok := holeMap[c];!ok{
                holeMap[c] = &node{
                    char: c,
                    next: make(map[*node]bool,0),
                }
            }
        }
    }

    // 建图的边
    for i:=0;i<len(words)-1;i++{
        a,b := diff(words[i],words[i+1])

        if a == b{
            continue
        }
        ptr := holeMap[a]
        nptr := holeMap[b]
        ptr.next[nptr] = true
    }

    for w,ptr := range holeMap{
        fmt.Println(string(w))
        ss := ""

        for pntr := range ptr.next{
            ss += string(pntr.char)
            ss += " "
        }
        fmt.Println(ss)
        fmt.Println("------")
    }

    IsOk := make(map[byte]bool)
    
    // 计算出入度
    for s,ptr := range holeMap{
        for c,_ := range ptr.next{
            ptr.outDim++
            nptr := holeMap[c.char]
            nptr.inDim++
        } 
        if ptr.outDim == 0{
            IsOk[s] = true
        }
    }

    // 检测是否有环
    isOk := make(map[byte]bool)
    path := make(map[byte]bool)
    r := true
    for c,_ := range holeMap{
        hasCircle(c,&holeMap,&isOk,&path,&r)
        if !r{
            return ""
        }
    }

    uu := ""
    // 通过出入度计算
    for q:=0;q<len(holeMap);q++{
        for c,ptr := range holeMap{
            if ptr.outDim == 0{
                uu = string(c) + uu
                ptr.outDim = -1

                for _,nptr := range holeMap{
                    if nptr.next[ptr]{
                        nptr.outDim -= 1
                    }
                }
                break
            }
        }
    }
    fmt.Println(holeMap) 
    return uu
}

func main(){
	words := []string{"wrt","wrf","er","ett","rftt"}
	fmt.Println(alienOrder(words))
	fmt.Println("WEFWES")
}