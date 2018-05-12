const $help = $('.help-text--wrapper');
const $container = $('.container');
const parseInput = (input) => {
  return input.split(' ');
}

$('.close').on('click', () => {
  $help.hide();
  $container.removeClass('helping');
});

$('.help').on('click', () => {
  $help.show();
  $container.addClass('helping');
});

let nickname = {};

$.get('data/nickname.json')
  .done(data => {
    for (let key of Object.keys(data)) {
      const arr = data[key];
      for (let name of arr) {
        nickname[name] = key;
      }
    }
  });

$.get('data/data.json')
  .done(stats => {
    const $input = $('#main');
    $input.change(() => {
      const result = parseInput($input.val().trim());
      let index = 0;
      let redTotal = 0;
      let blueTotal = 0;
      for (text of result) {
        index ++;
        const key = nickname[text];
        const stat = stats[key];
        if (stat) {
          const win = stat.winning;
          const lose = stat.losing;
          const sum = win + lose;

          const result = `${win}/${lose}`;
          const winp = parseInt((win / sum) * 100) + '%';
          const score = 6.31 * win / sum + 2.08 * sum / 2090;
          
          const $target = $(`.result-${index}`);
          const avatar = `../resources/pixyys/${key}.png`;
          const placeholder = '../resources/pixyys/御行达摩.png'
          $target.find('.name').text(key);
          $.ajax({
            url: avatar
          }).done(function(data, txt, xhr) {
            if (xhr.status === 200) {
              $target.find('.avatar').css('background-image', `url('${avatar}')`);
            } else {
              $target.find('.avatar').css('background-image', `url(${placeholder})`);
            }
          }).fail(function() {
            $target.find('.avatar').css('background-image', `url(${placeholder})`);
          });
          $target.find('.win-lose').text(result);
          $target.find('.win-percentage').text(winp);
          $target.find('.score').text(parseInt(score * 100) / 100);
          if (index <= 5) {
            redTotal += score;
          } else {
            blueTotal += score;
          }
        }
      }
      $('.result-wrapper--red .overview').text(parseInt(redTotal / 5 * 100) / 100);
      $('.result-wrapper--blue .overview').text(parseInt(blueTotal / 5 * 100) / 100);
    });
  });