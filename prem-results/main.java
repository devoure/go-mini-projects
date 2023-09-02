import java.util.Scanner;
import java.util.Arrays;
import java.util.InputMismatchException;

public class main{
  public static void main(String[] args){
    int[] blueScores = new int[5];
    int[] redScores = new int[5];

    System.out.println("Get results for red team");
    redScores = addScores(redScores);

    System.out.println("Get results for blue team");
    blueScores = addScores(blueScores);

    System.out.println("Scores for red and blue team");
    System.out.println(Arrays.toString(blueScores));
    System.out.println(Arrays.toString(redScores));

    int[] leagueResults;
    leagueResults = setResults(redScores, blueScores);

    System.out.println("Final league results");
    System.out.println("Red : " + leagueResults[0]);
    System.out.println("Blue : " + leagueResults[1]);
    if (leagueResults[0] > leagueResults[1]){
      System.out.println("Red wins");
    }
    else if (leagueResults[0] < leagueResults[1]){
      System.out.println("Blue wins");
    }
    else{
      System.out.println("No winner");
    }
  }
  public static int[] addScores(int[] teamScores){
    Scanner scanner = new Scanner(System.in);

    for (int i = 0; i < teamScores.length; i++){
      System.out.print("Enter team score for game " + (i + 1) + " : ");
      boolean isScoreSet = false;
      System.out.println("");
      while (isScoreSet == false){
        try{
          teamScores[i] = scanner.nextInt();
          isScoreSet = true;
        }
        catch(InputMismatchException e){
          System.out.println("Invalid Input, ERROR!");
          System.out.println("Try again");
          // remove input from buffer
          // replace with the next input
          scanner.nextLine();
        }
      }
    }
    return teamScores;
  }
  public static int[] setResults(int[] redScores, int[] blueScores){
    int redPoints = 0;
    int bluePoints = 0;
    // 3 for win
    // 1 for draw
    // 0 for loss
    for (int i = 0; i < 5; i++){
      // Red wins
      if (redScores[i] > blueScores[i]){
        redPoints = redPoints + 3;
        System.out.println("Red wins game " + (i + 1));
      }
      // blue wins
      else if (redScores[i] < blueScores[i]){
        bluePoints = bluePoints + 3;
        System.out.println("Blue wins game " + (i + 1));
      }
      // Draw
      else{
        redPoints = redPoints + 1;
        bluePoints = bluePoints + 1;
        System.out.println("Game " + (i + 1) + "ends in draw");
      }
    }
    int[] leagueResults = new int[2];
    leagueResults[0] = redPoints;
    leagueResults[1] = bluePoints;
    return leagueResults;
  }
}
