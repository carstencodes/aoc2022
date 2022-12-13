using System.ComponentModel.DataAnnotations;
using System.Text;

namespace aoc2022.day12;

internal class Parser {
    private const char StartCharacter = 'S';
    private const char EndCharacter = 'E';
    private const char Lowest = 'a';
    private const char Highest = 'z';

    internal Matrix Parse(string data) {
        Dictionary<Point, char> characters = new();

        Point? start = null;
        Point? end = null;

        List<Point> lowestPoints = new();

        uint rowId = 0;
        foreach (string line in data.Split(Environment.NewLine, StringSplitOptions.None)) {
            uint colId = 0;
            foreach (char value in line) {
                Point p = new (rowId, colId);
                char item = value;
                if (value == Parser.StartCharacter) {
                    start = p;
                    item = Parser.Lowest;
                }
                else if (value == Parser.EndCharacter) {
                    end = p;
                    item = Parser.Highest;
                }

                if (item == Parser.Lowest){
                    lowestPoints.Add(p);
                }

                characters.Add(p, item);

                colId++;
            }
            rowId++;
        }

        return Matrix.CreateNew(characters, start, end, lowestPoints.ToArray());
    }
}

internal readonly struct Point : IEquatable<Point> {
    internal readonly uint RowId = 0;
    internal readonly uint ColId = 0;

    internal Point(uint rowId, uint colId) {
        RowId=rowId;
        ColId=colId;
    }

    public override string ToString()
    {
        return $"({this.RowId}, {this.ColId})";
    }

    public override int GetHashCode()
    {
        int result = base.GetHashCode();
        return HashCode.Combine(result, this.RowId, this.ColId);
    }

    public bool Equals(Point other)
    {
        return this.RowId == other.RowId && this.ColId == other.ColId;
    }

    internal IEnumerable<Point> GetNeighbors() {
        // Up
        if (this.RowId > 0) {
            yield return new Point(this.RowId - 1u, this.ColId);
        }
        // Down
        yield return new Point(this.RowId + 1, this.ColId);
        // Left
        if (this.ColId > 0) {
            yield return new Point(this.RowId, this.ColId - 1u);
        }
        // Right
        yield return new Point(this.RowId, this.ColId + 1);
    }
}

internal sealed class Matrix {
    private class DirectionComparer : IComparer<char>
    {
        public int Compare(char from, char to)
        {
            if (from == to) {
                return 0;
            }

            if (from > to) {
                return 1;
            }

            if (to == from + 1) {
                return 1;
            }

            return -1;
        }
    }

    private readonly Dictionary<Point, char> pointToCharacters = new();

    private readonly IComparer<char> directionComparer;

    public Point Start { get; init; }
    public Point End { get; init; }

    public IReadOnlyCollection<Point> LowestPoints{ get; init;}

    private Matrix(IReadOnlyDictionary<Point, char> points)
    {
        this.pointToCharacters = new(points);
        this.directionComparer = new DirectionComparer();
    }

    internal static Matrix CreateNew(IReadOnlyDictionary<Point, char> points, Point? start, Point? end, params Point[] lowestPoints) {
        return new Matrix(points) {
            Start = start ?? throw new ArgumentNullException(nameof(start)),
            End = end ?? throw new ArgumentNullException(nameof(end)),
            LowestPoints = lowestPoints ?? Array.Empty<Point>()
        };
    }

    public char? this[uint row, uint col] {
        get {
            Point p = new (row, col);
            return this[p];
        }
    }

    public char? this[Point p] {
        get {
            if (!this.pointToCharacters.TryGetValue(p, out char value)) {
                return null;
            }

            return value;
        }
    }

    public override string ToString()
    {
        StringBuilder buffer = new();

        bool line_found = true;
        uint line_ctr = 0;
        while (line_found) {
            bool col_found;
            uint col_ctr = 0;
            do
            {
                Point current = new(line_ctr, col_ctr);
                col_found = this.pointToCharacters.TryGetValue(current, out char value);
                if (col_found) {
                    buffer.Append(value);
                    col_ctr ++;
                }
            } while (col_found);
            buffer.AppendLine();
            line_ctr++;
            line_found = this.pointToCharacters.ContainsKey(new Point(line_ctr, 0));
        }

        return buffer.ToString();
    }

    public SpanningTree ToTree(Point startPosition) {
        SpanningTree root = SpanningTree.CreateRoot(startPosition, this[startPosition]);
        Queue<Node> queue = new();
        HashSet<Point> visitedNodes = new();
        List<Point> traversedNodes = new();

        // Use BFS
        queue.Enqueue(root);
        visitedNodes.Add(root.Position);
        traversedNodes.Add(root.Position);

        while (queue.Count > 0) {
            Node node = queue.Dequeue();
            foreach(Point point in node.Position.GetNeighbors()) {
            if (visitedNodes.Contains(point)){
                continue;
            }

            char? value = this[point];
                if (value.HasValue && this.directionComparer.Compare(node.Value, value.Value) >= 0)
                {
                    Node newNode = node.AddNew(point, value.Value);
                    if (!point.Equals(this.End))
                    {
                        visitedNodes.Add(point);
                        traversedNodes.Add(point);
                        queue.Enqueue(newNode);
                    }
                    else
                    {
                        root.AddEndNode(newNode);
                    }
                }
            }
        }

        return root;
    }

    public SpanningTree ToTree() {
        return this.ToTree(this.Start);
    }

    public IEnumerable<SpanningTree> GetAllSpanningTreesFromLowestPoints() {
        foreach (Point p in this.LowestPoints) {
            yield return this.ToTree(p);
        }
    }
}

internal class SpanningTree : Node {
    private readonly List<Node> endNodes = new();

    internal IReadOnlyCollection<Node> EndingNodes => this.endNodes;

    internal static SpanningTree CreateRoot(Point position, char? value) {
        return new SpanningTree{
            Predecessor = null,
            Position = position,
            Value = value ?? throw new ArgumentNullException(nameof(value))
        };
    }

    internal void AddEndNode(Node node) {
        this.endNodes.Add(node);
    }
}

internal class Node : List<Node> {
    private readonly HashSet<Point> predecessors = new();

    internal Node? Predecessor {get; init;}

    internal char Value {get; init;}

    internal Point Position { get; init;}

    public override int GetHashCode()
    {
        return HashCode.Combine(this.Predecessor?.GetHashCode() ?? 0, this.Value, this.Position);
    }

    internal uint Weight {
        get {
            this.FillPredecessors();
            return Convert.ToUInt32(this.predecessors.Count);
        }
    }

    public override string ToString()
    {
        return $"{this.Predecessor?.ToString() ?? string.Empty} -> {this.Position}[{this.Value}]";
    }

    internal Node AddNew(Point position, char value) {
        Node newNode = new () {
            Predecessor = this,
            Position = position,
            Value = value
        };
        this.Add(newNode);
        return newNode;
    }

    internal bool IsSuccessorOf(Point p) {
        this.FillPredecessors();
        return this.predecessors.Contains(p);
    }

    private void FillPredecessors() {
        if (!this.predecessors.Any()) {
            Node? pre = this.Predecessor;
            while (pre != null) {
                this.predecessors.Add(pre.Position);
                pre = pre.Predecessor;
            }
        }
    }
}