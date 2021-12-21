using System;
namespace generics
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine(Adder.add<int>(3, 4));
            Console.WriteLine(Adder.add<double>(3.0, 4.0));
            Console.WriteLine(Adder.add<string>("3", "7"));
        }
    }
    public static class Adder
    {
        public static T add<T>(T a, T b)
        {
            T c = a + b;
            return c;
        }
    }
}

