import logging

import pandas as pd
import matplotlib.pyplot as plt


def load_benchmark_data(filename):
    try:
        data = pd.read_csv(filename)
        return data
    except Exception as e:
        print(f"Error loading CSV file: {e}")
        return None


def plot_performance_by_size(data):
    plt.figure(figsize=(12, 6))
    for solver in data['Solver'].unique():
        solver_data = data[data['Solver'] == solver]
        plt.plot(solver_data['Size'], solver_data['TimePerOperation'], label=solver, marker='o')

    plt.title("Performance Comparison by Matrix Size")
    plt.xlabel("Matrix Size")
    plt.ylabel("Time Per Operation (s)")
    plt.yscale('log')
    plt.legend()
    plt.grid(True)
    plt.show()


def main():
#     filename = "benchmarks_backup.csv"
    filename = "benchmarks.csv"
    try:
        data = pd.read_csv(filename)
    except Exception:
        logging.exception(f'Unable to load file {filename}!')
    else:
        print(data.head())
        plot_performance_by_size(data)


if __name__ == "__main__":
    main()
