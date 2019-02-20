package hashcode;

import java.io.FileNotFoundException;
import java.io.IOException;
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.Writer;
import java.io.BufferedWriter;
import java.io.OutputStreamWriter;
import java.io.FileOutputStream;
import java.util.ArrayList;
import java.util.List;


public class Main {
    static int row = 0;
    static int col = 0;
    static int minIngredients = 0;
    static int maxCells = 0;

    static char[][]board;

    public static void main(String[] args) throws IOException {
        String fileName  = "d_big";
        readFile(fileName + ".in");
        String outputFileName = fileName + ".out";
        getSlices(outputFileName);



    }

    public static void readFile(String filename) throws FileNotFoundException, IOException {
        BufferedReader br = new BufferedReader(new FileReader(filename));
        String line = "";
        line = br.readLine();
        String[] info = line.split(" ");
        row = Integer.parseInt(info[0]);
        col = Integer.parseInt(info[1]);
        minIngredients = Integer.parseInt(info[2]);
        maxCells = Integer.parseInt(info[3]);

        createBoard(row,col);

        line = br.readLine();
        int j = 0;

        while (line != null) {
            int i = 0;
            while(i < col && j < row) {
                    board[j][i] = line.charAt(i);
                    i++;

            }
            j++;
            line = br.readLine();
        }
    }

    public static void createBoard(int row, int col) {
        board = new char[row][col];
    }

    public static void getSlices(String outputFileName) throws IOException {
        List<List<Integer>> result = new ArrayList<>();

        for (int c=0; c< col; c++) {
            int beg = 0;
            int end = 0;
            int mush = 0;
            int tom = 0;

            while(end < row) {
                if (board[end][c] == 'M') {
                    mush++;
                } else if (board[end][c] == 'T') {
                    tom++;
                }

                end++;

                if (end - beg > maxCells) {
                    if (board[beg][c] == 'M') {
                        mush--;
                    } else if (board[beg][c] == 'T') {
                        tom--;
                    }
                    beg++;
                }


                if (end - beg <= maxCells && mush >= minIngredients && tom >= minIngredients) {
                    List<Integer> list = new ArrayList<>();
                    list.add(beg);
                    list.add(c);
                    list.add(end-1);
                    list.add(c);

                    beg = end;
                    tom = 0;
                    mush = 0;


                    result.add(list);
                }
            }
        }
        createOutput(result, outputFileName);
    }

    public static void createOutput(List<List<Integer>> result, String outputFileName) throws IOException {
        try (Writer writer = new BufferedWriter(new OutputStreamWriter(
                new FileOutputStream(outputFileName), "utf-8"))) {
            writer.write(result.size()+" ");

            for (List<Integer> list : result) {
                ((BufferedWriter) writer).newLine();
                writer.write(list.get(0) + " "  + list.get(1) + " " + list.get(2) + " " + list.get(3));
            }
        }
    }



}
