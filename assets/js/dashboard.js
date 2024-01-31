import Chart from "chart.js/auto";

// TODO: Make dynamic data
(async () => {
  const canvas = document.getElementById('spending');

  const data = {
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
  const centerLabel = {
    id: 'centerLabel',
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
  new Chart(
    canvas,
    {
      type: 'doughnut',
      data: data,
      plugins: [centerLabel],
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
      },
    }
  )

})()
