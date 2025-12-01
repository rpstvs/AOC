import java.io.File;                  // Import the File class
import java.util.ArrayList;
import java.util.List; // Import this class to handle errors
import java.util.Scanner; 

public class Solution{
    

    public static void main(String[] args) {
        List<String> result = ReadFile("input.txt");
        int zeros = FindPassword(result);
        System.out.println(zeros);
    }

    public static List<String> ReadFile(String filename){
        File input = new File(filename);
        List<String> result = new ArrayList<>();
        try(Scanner reader = new Scanner(input)) {

            while(reader.hasNextLine()){
                result.add(reader.nextLine());
            }
            
        } catch (Exception e) {
        }

        return result;
    }

    public static int FindPassword(List<String> Steps){
       
        int zeros = 0;
        int currentState = 50;
        int previousState = 50;
        for (String Step : Steps) {
            
            int stepInt = Integer.parseInt(Step.substring(1));
            int turns = stepInt / 100;
            int remainder = stepInt % 100;
            zeros += turns;
            //System.out.println("Step: " + Step + " Turns: "+ turns + " remainder " + remainder + " " + currentState + " " + zeros);
            switch (Step.charAt(0)) {
                case 'L':
                    currentState -= remainder;
                    if (currentState < 0) {
                        currentState = 100 + currentState;
                        if (currentState != 0 && previousState != 0){
                            zeros++;
                        }
                    }
                    break;
                 case 'R':
                    currentState += remainder;
                    if (currentState > 99) {
                        currentState = currentState - 100;
                        if (currentState != 0 && previousState != 0){
                            zeros++;
                        }
                        
                    }
                    break;
            }
            //System.out.println("The dial is rotated " +  Step + "to point at " + currentState);
            if (currentState == 0){
                zeros++;
            }
            
            previousState = currentState;
            //System.out.println("Adding Turns :" + turns + " to password " + zeros);
        }
        
        return zeros;
    }
}