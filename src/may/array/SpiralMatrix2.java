package may.array;

// https://leetcode-cn.com/problems/spiral-matrix-ii/
public class SpiralMatrix2{
    public int[][] generateMatrix(int n) {
        var rtn = new int[n][n];
        int is = 0;
        int js = 1;
        int i = 0, j = 0;
        int c = 1;
        while(c <= n * n){
            rtn[i][j] = c++;
            if(js != 0 && (j+js == n || j + js == -1 || rtn[i][j+js] != 0)){
                is = js;
                js = 0;
            }else if(is != 0 && (i + is == n || i + is == -1 || rtn[i+is][j] != 0)){
                js = -is;
                is = 0;
            }
            i += is;
            j += js;
        }
        return rtn;
    }

    public static void main(String[] args) {
        var t = new  SpiralMatrix2();
        var arr = t.generateMatrix(3);
        for(var c : arr){
            for(var v : c){
                System.out.print(v + " ");
            }
            System.out.println();
        }
    }
}