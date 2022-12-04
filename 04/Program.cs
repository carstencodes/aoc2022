using aoc2022.task04;

const string AllSections = Data.Puzzle;

string[] sections = AllSections.Split(Environment.NewLine, StringSplitOptions.RemoveEmptyEntries);

int numberOfOverlappingSections = 0;
int numberOfAnyOverlappingSections = 0;
foreach (string section in sections)
{
    string[] sectionsPerElf = section.Split(",", 2, StringSplitOptions.TrimEntries);

    string[] firstSectionBounds = sectionsPerElf[0].Split("-", 2);
    string[] secondSectionBounds = sectionsPerElf[1].Split("-", 2);


    (int firstLowerBound, int firstUpperBound) = (
        Int32.Parse(firstSectionBounds[0]),
        Int32.Parse(firstSectionBounds[1])
    );

    (int secondLowerBound, int secondUpperBound) = (
        Int32.Parse(secondSectionBounds[0]),
        Int32.Parse(secondSectionBounds[1])
    );

    ISet<int> firstSection = Enumerable.Range(firstLowerBound, firstUpperBound - firstLowerBound + 1).ToHashSet();
    ISet<int> secondSection = Enumerable.Range(secondLowerBound, secondUpperBound - secondLowerBound + 1).ToHashSet();

    if (firstSection.All(i => secondSection.Contains(i)) || secondSection.All(i => firstSection.Contains(i)))
    {
        numberOfOverlappingSections++;
    }

    if (firstSection.Intersect(secondSection).Any())
    {
        numberOfAnyOverlappingSections++;
    }
}

Console.WriteLine(numberOfOverlappingSections);
Console.WriteLine(numberOfAnyOverlappingSections);