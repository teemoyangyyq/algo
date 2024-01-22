# algo
golang各种算法库

//  二元组生成优化，对于两个数组  【1,1,1】 【1,1,1】， 正常情况会生成 {（1,3）,（1,2) , (1,1) ,（2,3）,（2,2) , (2,1),（3,3）,（3,2) , (3,1)}
//  我们看如何优化， 对于第一个数组的第一个值1， 应该生成（1,3）,（1,2) , (1,1)，我们看，（1,3），（1,2），是不是多余生成了，
//  因为对于单调队列， 先生成（1,3）， 丢入队列，再生成（1,2），会取代（1,3）丢入队列，最后生成（1,1),会取代（1,2）丢入队列，
//  (1,3),(1,2)生成白白浪费资源，一开始就直接生成（1，1），所以我们一开始对于每一个横坐标，纵坐标是有一个上限，即刚好大于单调队列最后的一个y值，此时可以用二分求出上限
//  此时 （2,3）， （2，2）， （2,1） 此时用上面的上限优化，保证只生成（2,2）， （2,1） ，此时我们发现（2,1）和（1,1）重复了，怎么去掉重复的（2,1），只生成（2,2）
//  对于同一个值1来说，它有9个二元组，假设第一个1生成了（1,1），第二个，第三个1不用生成（2,1），（3,1）， 因为找最大公共序列，对于同一纵坐标1来说，只能出现一次，横坐标最小肯定最优
//  所以我们遍历可以记录相同值已使用的纵坐标，下次对于相同值，肯定横坐标已使用的就不能使用，这样保证重复值的二元组只会生成不超过重复值个数次，以此保证生成二元组个数不会超过两个数组长度之和，
//  所以对于两个重复元素数组【1,1,1】 【1,1,1】只能会生成{（1,1），（2,2），（3,3）}这三个二元组

type T struct {
    X int 
    Y int
    Pre *T
}

func GetLine(data1,data2 []int)  []int {
    hash1, hash2 := make(map[int][]int), make(map[int][]int)
    posMap := make(map[int]int)
    for i := 0; i < len(data1); i++ {
        hash1[data1[i]] = append(hash1[data1[i]],i)
    }
    for i := 0; i < len(data2); i++ {
        hash2[data2[i]] = append(hash2[data2[i]],i)
    }

    P := make([]T,0)
    flag := false
    for i := 0; i < len(data1); i++ {
        if len(hash1[data1[i]]) && len(hash2[data1[i]]) > 0 {   
            if !flag {
                temp := T{
                    X :i ,
                    Y :hash2[data1[i]][0]
                }
                P := append(P,temp) 
                flag = true
                break
            }
            for j := LargePos(i, P, hash2[data1[i]]); j >= SmallPos(i, posMap, P, hash2[data1[i]]); j-- {
                temp := T{
                    X :i ,
                    Y :hash2[data1[i]][j]
                }
                pos := FindPos(j,P)
                temp.Pre = P[pos] 
                P[pos] = temp 
            }
              
        }    
    }
}


//  对于二元组同一横坐标，纵坐标只需从比单调队列大的数开始，通过二分查找去掉多余的二元组
func LargePos(x, P, PHash []T)  (pos int, ok bool) {
    pos := Find(0, len(PHash)-1,P[len(P) -1 ].X,P[len(P) -1 ].Y,PHash)    
    return
} 

// 对于二元组同一横坐标，纵坐标只需从比上次同一值的最大纵坐标大，通过二分查找去掉多余的二元组
func SmallPos(x int,  posMap map[int]int, P, PHash []T)  (pos int, ok bool) {
    v := 0
    if ok, _ := posMap[x]; !ok {
        v = posMap[x]
    }
    pos := Find(0, len(PHash)-1,P[0].X, v, PHash) 
    posMap[x] = pos   
    return
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
func Find(left, right, x, y int, PHash []T) (int, bool) {
    if left  == right {
        if Big(PHash[left].X,PHash[left].Y, x, y) {
            return left
        }
        return left+1
    }
    mid := (left + right) >> 1
    if Small(PHash[mid].X,PHash[mid].Y, x, y) {
        if !Small(PHash[mid+1].X,PHash[mid+1].Y,x, y){
            return  mid+1
        }
        return Find(mid+1, right, x, y, P)
    }
    return Find(left, mid, x, y, P)

}

func FindPos()








