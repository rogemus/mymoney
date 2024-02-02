import Chart from "chart.js/auto";

const loadData = () => {
  const labels = ["1.02", "2.02", "3.03", "4.02", "5.02"];
  return {
    labels: labels,
    datasets: [
      {
        data: [20, 300, 500, 100, 3000],
      }
    ]
  };
}

export const monthFlowChart = () => {
  const canvas = document.getElementById('flow');
  const data = loadData();
  const options = {
    responsive: true,
    plugins: {
      legend: {
        display: false
      },
    },
  }

  new Chart(canvas, {
    type: 'line',
    data: data,
    options
  });
}

