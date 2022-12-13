using aoc2022.day12;

Parser p = new Parser();
Matrix m = p.Parse(Data.Puzzle);
Console.WriteLine(m.ToString());
SpanningTree root = m.ToTree();
Node? minimumNode = root.EndingNodes.AsParallel().MinBy(node => node.Weight);

Console.WriteLine(minimumNode!.Weight);
Console.WriteLine(minimumNode!.ToString());

minimumNode = m.GetAllSpanningTreesFromLowestPoints().AsParallel().SelectMany(spanningTree => spanningTree.EndingNodes).MinBy(node => node.Weight);

Console.WriteLine(minimumNode!.Weight);
Console.WriteLine(minimumNode!.ToString());
