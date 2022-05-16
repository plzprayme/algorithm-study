class Solution {
    
    private int[] dx = {0, 1, 1, 1, 0, -1, -1, -1};
    private int[] dy = {1, 1, 0, -1, -1, -1, 0, 1};
    
    private boolean[][] visited;
    
    private int N;
    
    public int shortestPathBinaryMatrix(int[][] grid) {
        N = grid.length;
        visited = new boolean[N][N];
        
        if (grid[0][0] == 1) return -1;
        int answer = bfs(grid);
        return answer;
    }
    
    private int bfs(int[][] grid) {
        Queue<Cell> q = new LinkedList<>();
        q.add(new Cell(0, 0, 0));
        visited[0][0] = true;
        
        while(!q.isEmpty()) {
            Cell p = q.poll();
            
            if (p.x == N-1 && p.y == N-1) return p.step + 1;
            
            for (int i = 0; i < 8; i++) {
                int nx = dx[i] + p.x;
                int ny = dy[i] + p.y;
                if (isOut(nx) || isOut(ny)) continue;
                if (grid[ny][nx] == 1) continue;
                if (visited[ny][nx]) continue;
                q.add(new Cell(nx, ny, p.step + 1));
                visited[ny][nx] = true;
            }
            
        }
        
        return -1;
    }
    
    private boolean isOut(int i) {
        return i >= N || i < 0;
    }
    
    static class Cell {
        int x, y, step;
        
        public Cell(int x, int y, int s) {
            this.x = x;
            this.y = y;
            this.step = s;
        } 
    }
}
