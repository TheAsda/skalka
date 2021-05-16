using System;
using System.Collections;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using Skalka.ConfigParser;
using Skalka.Configs;
using Skalka.Executor;
using Skalka.Helpers;
using Skalka.TaskRunner;

namespace Skalka
{
    public class Application
    {
        private readonly IConfigParser _parser;
        private readonly IExecutor _executor;

        public Application(IConfigParser parser, IExecutor executor)
        {
            _parser = parser;
            _executor = executor;
        }

        public void Run(Args args)
        {
            var config = _parser.ReadConfig<FlowConfig>(args.ConfigFile);

            _executor.Fill(config.Env, config.Jobs);
            _executor.Start();
            Console.WriteLine("Starting executor");
            _executor.Wait();
        }
    }
}