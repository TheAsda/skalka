using System;
using System.Collections.Generic;
using Skalka.Configs;
using Skalka.Helpers;
using Skalka.TaskRunner;

namespace Skalka.Queue
{
    public class Queue
    {
        public string Name { get; set; }
        public List<Task> Tasks { get; set; }
        private int _current;

        public Queue(string name)
        {
            Name = name;
            Tasks = new List<Task>();
            _current = 0;
        }

        public void Fill(List<string> env, List<Step> steps)
        {
            steps.ForEach(step =>
            {
                Tasks.Add(new Task(step.Name,
                    step.Run,
                    EnvironmentHelper.MergeEnvironments(env, step.Env),
                    step.Dir));
            });
        }

        public Task GetNext()
        {
            if (IsEmpty())
            {
                throw new Exception("Queue is empty");
            }

            return Tasks[_current++];
        }

        public bool IsEmpty() => _current >= Tasks.Count;
    }
}