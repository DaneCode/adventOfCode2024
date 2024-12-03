using System;
using System.Collections.Generic;
using System.IO;
using System.Text.RegularExpressions;

class Program
{
    static void Main()
    {
        // File path
        string filePath = "d3.txt";

        // Get the results
        List<int> multiplicationResults = ExtractAndMultiply(filePath);

        // Calculate the sum of the results
        int totalSum = 0;
        foreach (int result in multiplicationResults)
        {
            totalSum += result;
        }

        // Print the sum of the results
        Console.WriteLine("Sum of multiplication results:");
        Console.WriteLine(totalSum);
    }

    // Function to extract and multiply numbers
    static List<int> ExtractAndMultiply(string filename)
    {
        List<int> results = new List<int>();
        // Updated pattern to strictly match valid mul(X,Y) instructions
        Regex pattern = new Regex(@"mul\((\d{1,3}),(\d{1,3})\)");

        foreach (string line in File.ReadLines(filename))
        {
            MatchCollection matches = pattern.Matches(line);
            foreach (Match match in matches)
            {
                int n1 = int.Parse(match.Groups[1].Value);
                int n2 = int.Parse(match.Groups[2].Value);
                results.Add(n1 * n2);
            }
        }

        return results;
    }
}