using aoc2022.task03;

string value = Data.RealData;

const char MinValueLowerCase = 'a';
const char MinValueUpperCase = 'A';
const int OffsetLowerUpper = 26;

string[] rucksacks = value.Split(Environment.NewLine, StringSplitOptions.RemoveEmptyEntries);
ICollection<char> items = new List<char>();

foreach (string rucksack in rucksacks) {
    ISet<char> compartment1 = new HashSet<char>(), compartment2 = new HashSet<char>();

    int weight = rucksack.Length;
    if (weight % 2 != 0) {
        throw new InvalidOperationException($"Rucksack {rucksack} has an unbalanced weight of {weight}.");
    }

    compartment1.UnionWith(rucksack[..(weight / 2)]);
    compartment2.UnionWith(rucksack[(weight / 2)..]);

    char[] commonContents = compartment1.Intersect(compartment2).ToArray();
    if (commonContents.Length > 1) {
        throw new InvalidOperationException($"Rucksack {rucksack} has more than one common item {new string(commonContents)} in compartments {new string(compartment1.ToArray())} and {new string(compartment2.ToArray())}.");
    }

    items.Add(commonContents[0]);
}

if (rucksacks.Length % 3 != 0) {
    throw new InvalidOperationException("Failed to process data as it is not a multiple of three rucksacks");
}

int ctr = 0;
ICollection<char> groups = new List<char>();
while (ctr < rucksacks.Length) {
    ISet<char> group = new HashSet<char>();
    group.UnionWith(rucksacks[ctr]);
    ctr ++;
    group.IntersectWith(rucksacks[ctr]);
    ctr ++;
    group.IntersectWith(rucksacks[ctr]);
    ctr ++;

    groups.Add(group.Single());
}

foreach (var element in new [] { items, groups })
{
    Console.WriteLine(new string(element.ToArray()));

    int[] values = element.Select(i => 1 + (char.IsUpper(i)
        ? (int)(i - MinValueUpperCase) + OffsetLowerUpper
        : (int)(i - MinValueLowerCase))).ToArray();
    Console.WriteLine(string.Join(" - ", values));

    int sum = values.Sum();
    Console.WriteLine(sum);
}
