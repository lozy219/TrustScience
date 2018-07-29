let nicknames = {};
let filenames;
let statsArr;

$.get('data/nickname.json?_v=7')
  .done(data => {
    for (let key of Object.keys(data)) {
      const arr = data[key];
      for (let name of arr) {
        nicknames[name] = key;
      }
    }
  });

$.get('data/filename.json?_v=12')
  .done(data => {
    filenames = data;
    $.get('data/data.json?_v=7')
      .done(stats => {
        statsArr = Object.entries(stats);
        for (let entry of statsArr) {
          const entryData = entry[1];
          entryData.sum = entryData.losing + entryData.winning;
          entryData.winP = entryData.winning / entryData.sum;
        }
        statsArr.sort((a, b) => {
          let aData = a[1];
          let bData = b[1];
          if (aData.winP < bData.winP) {
            return 1;
          } else if (aData.winP > bData.winP) {
            return -1;
          } else if (aData.sum > bData) {
            return 1;
          } else {
            return -1;
          }
        });
        for (let entry of statsArr) {
          const [name, stat] = entry;
          const filename = `resources/pixyys/${filenames[name]}.png?_v=1`;
          const {winning, losing, sum, winP} = stat;
          $('.result .col').append('<div class=row></div>');
          $('.avatar .row').last().html(`<img src="${filename}">`);
          $('.win .row').last().text(winning);
          $('.lose .row').last().text(losing);
          $('.sum .row').last().text(sum);
          $('.winp .row').last().text(parseInt((winning / sum) * 100) + '%');
        }
      });
  });
