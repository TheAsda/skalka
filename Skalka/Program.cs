using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.IO;
using System.Threading.Tasks;
using Autofac;
using CommandLine;
using Skalka.ConfigParser;
using Skalka.Executor;
using Skalka.IO;
using Skalka.Queue;
using Skalka.TaskRunner;
using IContainer = Autofac.IContainer;

namespace Skalka
{
    static class Program
    {
        private static void Main(string[] args)
        {
            Parser.Default.ParseArguments<Args>(args).WithParsed((parsedArgs) =>
            {
                InitializeContainer().Resolve<Application>().Run(parsedArgs);
            }).WithNotParsed(errors =>
            {
                foreach (var error in errors)
                {
                    Console.Error.WriteLine(error);
                }
            });
        }

        private static IContainer InitializeContainer()
        {
            var builder = new ContainerBuilder();

            builder.RegisterType<PathReader>().As<IPathReader>();
            builder.RegisterType<ConfigParser.ConfigParser>().As<IConfigParser>();
            builder.RegisterType<TaskRunner.TaskRunner>().As<ITaskRunner>();
            builder.RegisterType<QueueProcessorFactory>().As<IQueueProcessorFactory>();
            builder.RegisterType<Executor.Executor>().As<IExecutor>();
            builder.RegisterType<Application>();

            return builder.Build();
        }
    }
}