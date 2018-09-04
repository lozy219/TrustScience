const $help = $('.help-text--wrapper');
const $container = $('.container');

const toPercentage = value => {
  return isNaN(value) ? '???' : parseInt(value * 10000) / 100 + '%';
}

const parseInput = input => {
  return input.split(' ').filter(str => str.length > 0);
}

const clearResult = $target => {
  $target.find('.name').text('');
  $target.find('.avatar').css('background-image', 'none');
  $target.find('.win-lose').text('');
  $target.find('.win-percentage').text('');
  $target.find('.score').text('');
}

const uploadImage = event => {
  $input = $('#main');
  $input.val('载入中...');
  $.ajax({
    url: 'http://uygnim.com:8734/match',
    // url: 'http://localhost:8734/match',
    method: 'POST',
    data: new FormData($('#upload')[0]),
    processData: false,
    contentType: false
  }).done(data => {
    var result = data.join(' ').trim();
    if (result) {
      $input.val(result).change();
    } else {
      $input.val('载入失败').change();
    }
  }).fail(() => {
    $input.val('载入失败').change();
  });
};

$('.close').on('click', () => {
  $help.hide();
  $container.removeClass('helping');
});

$('.help').on('click', () => {
  $help.show();
  $container.addClass('helping');
});

$('.rank').on('click', () => {
  window.open('http://uygnim.com/yys/frontend/rank.html');
});

$('.github').on('click', () => {
  window.open('https://github.com/lozy219/TrustScience');
});

$('.photo').on('click', () => {
  $('#match').click();
});

let nicknames = {};
let filenames;

$.get('frontend/data/nickname.json?_v=guiqie')
  .done(data => {
    for (let key of Object.keys(data)) {
      const arr = data[key];
      for (let name of arr) {
        nicknames[name] = key;
      }
    }
  });

$.get('frontend/data/filename.json?_v=guiqie')
  .done(data => {
    filenames = data;
    $.get('frontend/data/weightedScore.json?_v=guiqie')
      .done(scores => {
        $.get('frontend/data/data.json?_v=guiqie')
          .done(content => {
            const stats = content['A'];
            const $input = $('#main');
            $input.bind('change keyup input', () => {
              var result = parseInput($input.val().trim());
              var redTotal = 0;
              var blueTotal = 0;
              for (let index = 0; index < 10; index ++) {
                const $target = $(`.result-${index + 1}`);

                if (index >= result.length) {
                  clearResult($target);
                  continue;
                }

                const text = result[index];
                const key = nicknames[text];
                const stat = stats[key];
                const filename = filenames[key];

                if (!stat) {
                  clearResult($target);
                  continue;
                }

                const win = stat.winning;
                const lose = stat.losing;
                const sum = win + lose;

                const history = `${win}/${lose}`;
                let winp = parseInt((win / sum) * 100) + '%';
                if (sum === 0) {
                  winp = '0%';
                }

                const avatar = `frontend/resources/pixyys/${filename}.png?_v=2`;
                // const placeholder = 'frontend/resources/pixyys/yxdm.png'
                $target.find('.name').text(key);
                $target.find('.avatar').css('background-image', `url('${avatar}')`);
                $target.find('.win-lose').text(history);
                $target.find('.win-percentage').text(winp);
                if (index < 5) {
                  redTotal += scores[key];
                } else {
                  blueTotal += scores[key];
                }
              }
              const scoreSum = redTotal + blueTotal;
              $('.result-wrapper--red .overview').text(toPercentage(redTotal / scoreSum));
              $('.result-wrapper--blue .overview').text(toPercentage(blueTotal / scoreSum));
            });
          });
      });
  });
