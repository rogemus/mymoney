import Chart, { ChartConfiguration } from "chart.js/auto";

const loadData = () => {
  const labels = ["1.02", "2.02", "3.03", "4.02", "5.02", "6.02", "7.02", "8.02", "9.02"];
  return {
    labels: labels,
    datasets: [
      {
        data: [20, 300, 500, 100, 3000, 400, 801, 1400],
      }
    ]
  };
}

export const monthFlowChart = () => {
  const canvas = document.getElementById('flow') as HTMLCanvasElement;

  const chartConfig: ChartConfiguration<'line'> = {
    type: 'line',
    data: loadData(),
    options: {
      plugins: {
        legend: {
          display: false
        }
      }
    }
  }

  new Chart(canvas, chartConfig);
}

