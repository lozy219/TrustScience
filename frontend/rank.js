let nicknames = {};
let filenames;
let statsArr = [];

// function declarations

const sortStats = () => {
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
};

const clearTable = () => {
  // [TODO: clear the table content for re-render]
};

const renderTable = () => {
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
}

const _refresh = () => {
  sortStats();
  clearTable();
  renderTable();
};


// event handlers

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
        // populate stas array
        statsArr = Object.entries(stats);
        for (let entry of statsArr) {
          const entryData = entry[1];
          entryData.sum = entryData.losing + entryData.winning;
          entryData.winP = entryData.winning / entryData.sum;
        }
        // refresh UI
        _refresh();
      });
  });

$('.sort-icon').click(function() {
  let sort = $(this).attr('sort');
  let col = $(this).attr('col');

  sort_column(sort, col);
})

$('.col-header').click(function() {
  let col = $(this).attr('col');
  let sort = $('#' + col + '-sort-icon').attr('sort');

  sort_column(sort, col);
})

function sort_column(sort, col) {
  // change all to default first
  $('.sort-icon').attr('src', 'resources/icon/sort-solid.svg');
  $('.sort-icon').attr('sort', 'default');

  if (sort == 'default') {
    // change to desc order    
    $('#' + col + '-sort-icon').attr('src', 'resources/icon/sort-down-solid.svg');
    $('#' + col + '-sort-icon').attr('sort', 'desc');
  } else if (sort == 'desc') {
    // change to asc order
    $('#' + col + '-sort-icon').attr('src', 'resources/icon/sort-up-solid.svg');
    $('#' + col + '-sort-icon').attr('sort', 'asc');
  } else {
    // change to default order
    $('#' + col + '-sort-icon').attr('src', 'resources/icon/sort-solid.svg');
    $('#' + col + '-sort-icon').attr('sort', 'default');
  }
}

// $([TODO: table header selector]).on('click', event => {
//   [TODO: update sort mode]
//   _refresh();
// });

