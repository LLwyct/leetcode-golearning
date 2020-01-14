并查集求连通块数，n个连通块至少需要n-1条线相连，判断是否满足这个条件

var bc [100007]int;

// 自带路径压缩的find()
func find(x int) int {
    if x != bc[x] {
        bc[x] = find(bc[x])
    }
    return bc[x];
}

func makeConnected(n int, connections [][]int) int {
    var connectionNum = len(connections);
    for i:=0;i<n;i++ {
        bc[i] = i;
    }

    var blockNum = n;

    for i:=0;i<connectionNum;i++ {
        var a, b int = connections[i][0], connections[i][1];
        if find(a) != find(b) {
            bc[find(b)] = bc[find(a)];
            blockNum --;
        }
    }
    if connectionNum >= n - 1 {
        return blockNum - 1;
    } else {
        return -1;
    }
}
