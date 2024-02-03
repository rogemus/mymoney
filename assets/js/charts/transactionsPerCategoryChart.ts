import Chart, { ChartConfiguration, Plugin } from "chart.js/auto";

const centerLabelPlugin: Plugin<'doughnut'> = {
  id: 'centerLabelPlugin',
  afterDatasetDraw(chart, args) {
    const { ctx } = chart;
    const meta = args.meta;
    const amountFontSize = 32;
    const labelFontSize = 16;

    // amount
    const { x, y } = meta.data[0];
    const yAmount = y - amountFontSize / 2;

    ctx.textAlign = 'center';
    ctx.font = `bold ${amountFontSize}px sans-serif`;
    ctx.fillText(`$ ${meta.total}`, x, yAmount);
    ctx.save()
    ctx.restore();

    // label 
    ctx.textAlign = 'center';
    ctx.font = `bold ${labelFontSize}px sans-serif`;
    ctx.fillText('SPEND THIS MONTH', x, yAmount + amountFontSize);
    ctx.save()
    ctx.restore();
  }
}

const loadData = () => {
  return {
    labels: [
      'Red',
      'Blue',
      'Yellow'
    ],
    datasets: [{
      label: 'My First Dataset',
      data: [300.21, 50.14, 13.54],
      // backgroundColor: [
      //   'rgb(255, 99, 132)',
      //   'rgb(54, 162, 235)',
      //   'rgb(255, 205, 86)'
      // ],
      hoverOffset: 4
    }]
  };
}

export const transactionsPerCategoryChart = () => {
  const canvas = document.getElementById('spending') as HTMLCanvasElement;
  const chartConfig: ChartConfiguration<'doughnut'> = {
    type: 'doughnut',
    data: loadData(),
    plugins: [centerLabelPlugin],
    options: {
      plugins: {
        legend: {
          title: {
            display: true,
            padding: 5
          },
          position: 'bottom'
        }
      }
    }
  }
  new Chart(canvas, chartConfig);
}
