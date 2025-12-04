
import java.io.File;                  // Import the File class
import java.util.ArrayList;
import java.util.List; // Import this class to handle errors
import java.util.Scanner;

//read file
//separate the ranges
//parse the ranges
public class Solution {

    public static void main(String[] args) {
        List<String> result = ReadFile("input.txt");
        int Invalidids = getTotalJoltage(result);
        System.out.println(Invalidids);
    }

    public static List<String> ReadFile(String filename) {
        File input = new File(filename);
        List<String> result = new ArrayList<>();
        try (Scanner reader = new Scanner(input)) {

            while (reader.hasNextLine()) {
                result.add(reader.nextLine());
            }

        } catch (Exception e) {
        }

        return result;
    }

    public static int getTotalJoltage(List<String> Batteries) {

        int total = 0;
        int max = 0;
        for(var bat : Batteries){
            for (int i = 0; i < bat.length()-1; i++) {
                int n = bat.charAt(i) - '0';
                for (int j = i +1; j < bat.length(); j++) {
                    int k = bat.charAt(j) - '0';
                   int localmax = (n * 10) + k;

                    if (localmax > max){
                        max = localmax;
                    }

                }
                
                
            }
            total += max;
                max = 0;
           
        }
         return total;
    }
}