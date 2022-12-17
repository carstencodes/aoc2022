using aoc2022.day16;

Parser p = new Parser(Data.Sample);

ValveSet v = p.Parse();

Console.WriteLine(v);

Engine e = new Engine(v, 30, "AA");

long accumulated_max = e.Run();

Console.WriteLine(accumulated_max);