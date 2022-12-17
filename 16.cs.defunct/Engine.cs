namespace aoc2022.day16;

internal sealed class Engine {
    private readonly ValveSet valves;

    private readonly byte maxDuration;

    private readonly string beginAt;

    internal Engine(ValveSet valves, byte maxDuration, string beginAt) {
        this.valves = valves ?? throw new ArgumentNullException(nameof(valves));
        this.maxDuration = maxDuration;
        this.beginAt = beginAt ?? throw new ArgumentNullException(nameof(beginAt));
    }

    public long Run(){
        Valve start = this.valves[this.beginAt];
        IStep current = new MoveStep(start, this.maxDuration);
        IEnumerable<IStep> leafs = current.RunToEnd(this.valves);

        return leafs.AsParallel().Max(l => l.WayPathValue);
    }
}

internal interface IStep {

    IReadOnlyCollection<IStep> Predecessors { get; }

    long Value { get; }

    byte RemainingTimeStamps { get; }

    Valve Current { get; }

    long WayPathValue { get; }

    bool HasOpened(Valve valve);
}

internal abstract class StepBase : IStep {
    public IReadOnlyCollection<IStep> Predecessors { get; init; }

    public abstract long Value { get; }

    public byte RemainingTimeStamps { get; init; }

    public Valve Current { get; init; }

    public long WayPathValue { get; init; }

    protected IReadOnlyCollection<Valve> OpenedValves => this.openedValves;

    private readonly List<Valve> openedValves = new();

    protected StepBase(Valve current, byte remainingTimeStamps, params IStep[] predecessors) {
        this.Predecessors = predecessors ?? Array.Empty<IStep>();
        this.RemainingTimeStamps = remainingTimeStamps;
        this.Current = current;
        long currentWayPath = 0;
        IStep? recent = this.Predecessors.LastOrDefault();
        if (recent != null) {
            currentWayPath = recent.WayPathValue;
            if (recent is StepBase step) {
                this.openedValves.AddRange(step.openedValves);
            }
        }

        this.WayPathValue = currentWayPath + this.Value;
    }

    public bool HasOpened(Valve valve) {
        return this.openedValves.Contains(valve);
    }

    protected void AppendAsOpenedValve() {
        this.openedValves.Add(this.Current);
    }
}

internal sealed class MoveStep : StepBase
{
    internal MoveStep(Valve current, byte remainingTimeStamps, params IStep[] predecessors): base(current, remainingTimeStamps, predecessors) {
    }

    public sealed override long Value => 0;
}

internal sealed class OpenStep : StepBase
{
    internal OpenStep(Valve current, byte remainingTimeStamps, params IStep[] predecessors): base(current, remainingTimeStamps, predecessors) {
        this.AppendAsOpenedValve();
    }

    public sealed override long Value => this.RemainingTimeStamps * this.Current.MaxFlowRate;
}

internal static class StepExtensions {
    internal static IStep Move(this IStep step, Valve next) {
        List<IStep> predecessors = new(step.Predecessors)
        {
            step
        };
        return new MoveStep(next, Convert.ToByte(step.RemainingTimeStamps - 1), predecessors.ToArray());
    }

    internal static IStep Open(this IStep step) {
        List<IStep> predecessors = new(step.Predecessors)
        {
            step
        };
        return new OpenStep(step.Current, Convert.ToByte(step.RemainingTimeStamps - 1), predecessors.ToArray());
    }

    internal static IEnumerable<IStep> RunToEnd(this IStep start, ValveSet valves) {
        Queue<IStep> queue = new();
        queue.Enqueue(start);
        while (queue.Any()) {
            IStep step = queue.Dequeue();
            if (step.RemainingTimeStamps == 0) {
                yield return step;
            }
            else {
                foreach (IStep next in step.GetSuccessors(valves)) {
                    queue.Enqueue(next);
                }
            }
        }
    }

    private static IEnumerable<IStep> GetSuccessors(this IStep step, ValveSet valves) {
        if (step.RemainingTimeStamps == 0) {
            yield break;
        }

        foreach (string successor in step.Current.Successors) {
            Valve next = valves[successor];
            yield return step.Move(next);
        }
        
        if (step is not OpenStep && !step.Current.Stuck && !step.HasOpened(step.Current)) {
            yield return step.Open();
        }
    }
}