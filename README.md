func Adjust(node int, data []int) {
    left, right := node << 1 + 1, node << 1 + 2
    if left  >=  len(data) {
        return
    }
    maxPos := left
    if right < len(data) && data[right] > data[left]{
          maxPos = right
    }
    temp := data[maxPos]
    if data[maxPos] > data[node] {
        data[maxPos] = data[node]
        data[node]   = temp
        Adjust(maxPos, data)
    }
}


type T struct {
    X int 
    Y int
    Pre *T  // 反查指针
}

func GetLine(data1,data2 []int)  []int {
    hash := make(map[int][]int)
    posMap := make(map[int]int)
    // hash值坐标数组
    for i := 0; i < len(data2); i++ {
        hash[data2[i]] = append(hash[data2[i]], i)
    }
    // 单调二分队列
    P := make([]T,0)
    flag := false
    // 
    for i := 0; i < len(data1); i++ {
        if  len(hash[data1[i]]) > 0 { 
            // 单调数组为空时，直接放入单调数组，只会在单调队列初始化时调用 
            if !flag {
                temp := T{
                    X :i ,
                    Y :hash[data1[i]][0]
                }
                P := append(P,temp) 
                flag = true
                break
            }
            t := 0
            ok, t := posMap[i]
            if !ok {
                t = 0
            }
            // 生成二元组优化，保证二元组生成数量小于数组长度，上限
            for j := Min(LargePos(i, P, hash[data1[i]]), len(hash[data1[i]])-1); j >=0 ; j-- {
                temp := T{
                    X :i ,
                    Y :hash[data1[i]][j]
                }
                // 找到应该替换的二元组位置
                pos := FindPos(i,j,P)
                // 生成二元组优化，保证二元组生成数量小于数组长度，下限，如果相同值对应的子序列长度小于已存在的横纵坐标更小的子序列长度，结束
                if pos <= t {
                    break 
                }
                posMap[i] = pos
                
                
                // 对于替换保存记录，以便反查使用贪心搜索算法生成最长相同子序列
                temp.Pre = P[pos] 
                P[pos] = temp 
              
            }
              
        }    
    }
    // 生成最长公共子序列，待实现
    ResverseList(P)
}

1,2,1,1
2,1,1
/**
// 二元组生成优化，对于两个数组 【1,1,1】 【1,1,1】， 正常情况会生成 {（1,3）,（1,2) , (1,1) ,（2,3）,（2,2) , (2,1),（3,3）,（3,2) , (3,1)} 
// 我们看如何优化， 对于第一个数组的第一个值1， 应该生成（1,3）,（1,2) , (1,1)，我们看，（1,3），（1,2），是不是多余生成了， 
// 因为对于单调队列， 先生成（1,3）， 丢入队列，再生成（1,2），会取代（1,3）丢入队列，最后生成（1,1),会取代（1,2）丢入队列，
 // (1,3),(1,2)生成白白浪费资源，一开始就直接生成（1，1），所以我们一开始对于每一个横坐标，纵坐标是有一个上限，即刚好大于单调队列最后的一个y值，此时可以用二分求出上限 
// 此时 （2,3）， （2，2）， （2,1） 此时用上面的上限优化，保证只生成（2,2）， （2,1） ，此时我们发现（2,1）和（1,1）重复了，怎么去掉重复的（2,1），只生成（2,2）
 // 对于同一个值1来说，它有9个二元组，假设第一个1生成了（1,1），第二个，第三个1不用生成（2,1），（3,1）， 
 // 因为找最大公共序列，对于同一纵坐标1来说，如果两个横坐标的在单调队列的下标相同，(即长度相同，我们肯定取更小的横坐标），我们可以取更小的横坐标
 // 所以我们遍历可以记录相同值已使用的纵坐标，下次对于相同值，肯定横坐标已使用的就不能使用，这样保证想要重复使用元素必须有间隔，重复值使用一个元素，就会产生至少一个间隔数，只会生成不超过重复值个数次，以此保证生成二元组个数不会超过两个数组长度之和，
  // 所以对于两个重复元素数组【1,1,1】 【1,1,1】只能会生成{（1,1），（2,2），（3,3）}这三个二元组
**/
//  对于二元组同一横坐标，纵坐标只需从比单调队列大的数开始，通过二分查找去掉多余的二元组
func LargePos(x, P, PHash []T)  (pos int, ok bool) {
    pos := Find(0, len(PHash)-1,P[len(P) -1].X,P[len(P) -1].Y,PHash)    
    return
} 

func Min(x,y int) int{
    if x < y {
        return x
    }
    return y
}

func Max(x,y int) int{
    if x > y {
        return x
    }
    return y
}

                  
func Small(x1, y1, x2, y2 int) bool {
        return y1 < y2                                               
}
func Equal(x1, y1, x2, y2 int) bool {
    return y1 == y2                                               
}
func Big(x1, y1, x2, y2 int) bool {
        return y1 > y2                                               
}

// 找到比y大的hash数组第一个值下标
func Find(left, right, x, y int, PHash []T) (int, bool) {
    if left  == right {
        if Big(PHash[left].X,PHash[left].Y, x, y) {
            return left
        }
        return left+1
    }
    mid := (left + right) >> 1
    if !Small(x, y, PHash[mid].X, PHash[mid].Y) {
        if Small(PHash[mid+1].X, PHash[mid+1].Y, x, y){
            return  mid+1
        }
        return Find(mid+1, right, x, y, P)
    }
    return Find(left, mid, x, y, P)

}
// 查找单调队列位置，替换
func FindPos(){

}







