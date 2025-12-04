import java.io.File;                  // Import the File class
import java.util.ArrayList;
import java.util.List; // Import this class to handle errors
import java.util.Scanner;

//read file
//separate the ranges
//parse the ranges


public class Solution{
    

    public static void main(String[] args) {
        String result = ReadFile("input.txt");
        Long Invalidids = InvalidIdsThroughPattern(result);
        System.out.println(Invalidids);
    }

    public static String ReadFile(String filename){
        File input = new File(filename);
        List<String> result = new ArrayList<>();
        try(Scanner reader = new Scanner(input)) {

            while(reader.hasNextLine()){
                result.add(reader.nextLine());
            }
            
        } catch (Exception e) {
        }

        return result.get(0);
    }

    public static Long InvalidIdSum(String Ranges){
        
         String[] splitranges =Ranges.split(",");
        Long invalidIdSum = 0L;
         for (String range : splitranges){
            String[] splitRange = range.split("-");
            Long begin = Long.parseLong(splitRange[0]);
            Long end = Long.parseLong(splitRange[1]);
            Boolean invalidId = false;
            for (Long i = begin; i <= end; i++) {
                String value = i.toString();
                Integer midIndex = value.length() /2;
                
                if ((value.length() % 2) == 0){
                    // 123123
                   

                    String leftSide = value.substring(0,midIndex);
                    String rightSide = value.substring(midIndex);

                    invalidId = leftSide.equals(rightSide);
                }
                if (invalidId){
                    System.out.println("Invalid ID Detected " + i + " " + i.toString().length());
                    invalidIdSum += i;
                    invalidId = false;
                }
            }
         }

        
        return invalidIdSum;
    }

    public static Long InvalidIdsThroughPattern(String Ranges){
        
        String[] splitranges = Ranges.split(",");
        Long invalidIdSum = 0L;
         for (String range : splitranges){
            String[] splitRange = range.split("-");
            Long begin = Long.parseLong(splitRange[0]);
            Long end = Long.parseLong(splitRange[1]);
            Boolean hasPattern = false;
            for (Long i = begin; i <= end; i++) {
                String value = i.toString();
                Integer midIndex = value.length() /2;
                
                for (int j = 0; j < value.length() - 1; j++) {
                    var pattern = value.substring(0, j+1);
                     hasPattern = stringComposedEntirelyOfPattern(pattern, value);
                     if (hasPattern){
                            System.out.println("Invalid ID Detected " + i + " " + i.toString().length());
                            invalidIdSum += i;
                            break;
                        }
                }
                
            }
         }

        
        return invalidIdSum;
    }

   private static Boolean stringComposedEntirelyOfPattern(String pattern, String entire) {
        for (var i = 0; i < entire.length(); i += pattern.length()) {
            if (i + pattern.length() > entire.length()) {
                return false;
            }
            var target = entire.substring(i, i + pattern.length());
            if (!target.equals(pattern)) {
                return false;
            }
        }

        return true;
    }
}

