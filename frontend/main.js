const VERSION = 'citibank';
const $help = $('.help-text--wrapper');
const $container = $('.container');

// const host = '0.0.0.0';
const host = 'uygnim.com';
const wordMap = ['一', '两', '三', '四', '五', '六', '七', '八', '九', '十'];
let votingDisabled = false;

const disableVoting = () => {
  votingDisabled = true;
  $('.pointer').addClass('hidden');
  $('.hints').text('超谢谢你哦！');
};

const getCurrentTimeFrame = () => {
  return parseInt((new Date()).getTime() / 1000 / 60 / 5);
}

const vote = index => {
  if (votingDisabled) {
    return;
  }
  let selector = ''
  if (index === 0) {
    selector = '.previous-count-red';
  } else if (index === 1) {
    selector = '.previous-count-blue';
  } else {
    return;
  }

  $.get(`https://${host}:/api/report/${index}`)
    .done(data => {
      const count = data.count;
      $(selector).text(count);
      localStorage.setItem('ts_voted', getCurrentTimeFrame());
      disableVoting();
    });
}

const toPercentage = value => {
  return isNaN(value) ? '???' : parseInt(value * 10000) / 100 + '%';
}

const toStrike = value => {
  if (value == 0) {
    return '没出场';
  }
  const strikeWord = wordMap[Math.abs(value) - 1] || '好多';
  return value > 0 ? `${strikeWord}连胜` : `${strikeWord}连败`;
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
    url: `https://${host}/api/match`,
    method: 'POST',
    timeout: 200000,
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
  $help.find('.help-text').html(`
    这是一个根据1145场对弈竞猜记录制作的计算器。<br><br>
    上传截图或者手动输入双方十个式神名称后，会显示预测的胜率。<br><br>
    具体细节请看<a target="blank" href="https://bbs.ngacn.cc/read.php?tid=14044587">这个帖子</a>。
  `);
  $help.show();
  $container.addClass('helping');
});

const notesId = 'notes-1-read';
if (!localStorage.getItem(notesId)) {
  $('.notes').addClass('fresh');
}

$('.notes').on('click', () => {
  localStorage.setItem(notesId, true);
  $help.find('.help-text').html(`
    因为录入数据实在太麻烦了，现在页面会预加载别人已经上传过的阵容啦。<br><br>
    特别特别感谢八岐大熊的体验服数据！
  `);
  $help.show();
  $('.notes').removeClass('fresh');
  $container.addClass('helping');
});

$('.rank').on('click', () => {
  window.open('https://uygnim.com/yys/frontend/rank.html');
});

$('.github').on('click', () => {
  window.open('https://github.com/lozy219/TrustScience');
});

$('.photo').on('click', () => {
  $('#match').click();
});

let nicknames = {};
let filenames;

$.get(`frontend/data/nickname.json?_v=${VERSION}`)
  .done(data => {
    for (let key of Object.keys(data)) {
      const arr = data[key];
      for (let name of arr) {
        nicknames[name] = key;
      }
    }
  });

$.get(`frontend/data/filename.json?_v=${VERSION}`)
  .done(data => {
    filenames = data;
    $.get(`frontend/data/weightedScore.json?_v=${VERSION}`)
      .done(scores => {
        $.get(`frontend/data/data.json?_v=${VERSION}`)
          .done(content => {
            $.get(`frontend/data/strike.json?_v=${VERSION}`)
              .done(strike => {
                const stats = content['A'];
                stats['御行达摩'] = {
                  'losing': 0,
                  'winning': 0
                };
                scores['御行达摩'] = 50;
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
                    const key = nicknames[text] || '御行达摩';
                    const stat = stats[key];
                    const stk = strike[key] || 0;
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
                      winp = '50%';
                    }

                    const avatar = `frontend/resources/pixyys/${filename}.png?_v=3`;
                    const placeholder = 'frontend/resources/pixyys/yxdm.png'
                    $.get(avatar)
                      .done(() => {
                        $target.find('.avatar').css('background-image', `url('${avatar}')`);
                      })
                      .fail(() => {
                        $target.find('.avatar').css('background-image', `url('${placeholder}')`);
                      });
                    $target.find('.strike').text(toStrike(stk));
                    $target.find('.name').text(key);
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

                $.get(`https://${host}/api/result`)
                  .done(data => {
                    const current = data.current;
                    $('#main').val(current).change();

                    const record = data.previous.split(' ');
                    const result = data.result;
                    if (record.length === 10) {
                      $('.hidden').removeClass('hidden');
                      for (let i = 0; i < 10; i ++) {
                        const avatar = `frontend/resources/pixyys/${filenames[nicknames[record[i]]]}.png?_v=3`;
                        $(`.previous .avatar-${i + 1}`).css('background-image', `url('${avatar}')`);
                      }

                      $('.previous-count-red').text(result[0]);
                      $('.previous-count-blue').text(result[1]);

                      if (localStorage.getItem('ts_voted') == getCurrentTimeFrame()) {
                        disableVoting();
                      } else {
                        $('.previous-count-red').on('click', () => {
                          vote(0);
                        });
                        $('.previous-count-blue').on('click', () => {
                          vote(1);
                        });
                      }
                    }
                });
              });
          });
      });
  });
  
