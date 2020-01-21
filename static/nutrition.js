// code label container
let code_label = []

// start button opening first section
$('#start-button').click(function (e) {
    $(`#btn-collapse-intro`).trigger('click')
    $(`#btn-collapse-0`).trigger('click')
})

// renderResult renders details for each code segment
function renderResult(code) {
    const segments = code.split('_')
    let code_dict = {}
    $.getJSON(`${window.location.origin}/static/spectrum.json`, function (data) {
        $.each(data, function (key, entry) {
            $.each(entry.content, function (k, e) {
                code_dict[e.id] = {
                    "text": e.desc,
                    "title": entry.title
                }
            })
        })
        $.each(segments, function (k, e) {
            if (e) {
                $('#result').append(`<br><span>${code_dict[e].title}</span><br>`)
                $('#result').append(`<h3><span class="${getClassByContent(e)}">${e}</span> ${code_dict[e].text}<br></h3>`)
            }
        })
    })
}

// renderCardTitle renders a title section for one card
function renderCardTitle(parent, index, title) {
    $(parent).append(`
    <div class="card-header" id="heading-${index}">
        <h5 class="mb-0">
            <button id="btn-collapse-${index}" class="btn btn-link" data-toggle="collapse" data-target="#collapse-${index}" aria-expanded="false" aria-controls="collapse-${index}">
                ${title}
            </button>
        </h5>
    </div>`)
}

// renderCardBody renders a body section for one card
function renderCardBody(parent, index, body) {
    $(parent).append(`
    <div id="collapse-${index}" class="collapse" aria-labelledby="heading-${index}" data-parent="#accordion">
        <div class="card-body">
            ${body}
        </div>
    </div>`)
}

// renderSelectionMenu renders the selection menu for each category
function renderSelectionMenu() {
    $.getJSON(`${window.location.origin}/static/spectrum.json`, function (data) {
        $.each(data, function (index, entry) {
            $('#content').append($(`<div class="card" id="card-${index}"></div>`))
            renderCardTitle(`#card-${index}`, index, `<h3 style="display: inline;"><span class="badge badge-secondary" id="result-${index}"></span></h3> ${index+1}. ${entry.title}`)
            renderCardBody(`#card-${index}`, index, `<h4>${entry.desc}</h4><br><div class="list-group" id="card-body-${index}"></div>`)
            $.each(entry.content, function (k, e) {
                $(`#card-body-${index}`).append(`<button type="button" class="list-group-item list-group-item-action card-body-button-action" value="${e.id}" index="${index}">${e.desc}</button>`)
            })
        })

        $('.card-body-button-action').click(function (e) {
            let label = $(this).attr("value")
            let index = $(this).attr("index")
            code_label[index] = label
            $('#result-' + index).text(label)
            $('#result-' + index).attr('class', getClassByContent(label))
            $('#link').html(`<a href="facts/${getCodeUrlencoded()}"><img alt="code nutrition facts" src="badge/${getCodeUrlencoded()}"></img></a>`)
            
            $('#snippet-html').attr('value', `<a href="${window.location.origin}/facts/${getCodeUrlencoded()}"><img alt="code nutrition facts" src="${window.location.origin}/badge/${getCodeUrlencoded()}"></img></a>`)
            $('#snippet-md').attr('value', `[![nutrition facts](${window.location.origin}/badge/${getCodeUrlencoded()})](${window.location.origin}/facts/${getCodeUrlencoded()})`)

            $(`#btn-collapse-${parseInt(index) + 1}`).trigger('click')
            $(`#btn-collapse-${index}`).trigger('click')
            var element = document.querySelector(`#btn-collapse-${parseInt(index) + 1}`);
            element.scrollIntoView({
                behavior: 'smooth',
                block: 'end'
            });
        })

        lastitem = data.length
        $('#content').append($(`<div class="card" id="card-${lastitem}"></div>`))
        renderCardTitle(`#card-${lastitem}`, lastitem, `Result`)
        renderCardBody(`#card-${lastitem}`, lastitem, `<h4>Embed this Image</h4><br><div id="link"></div><br><h4>Copy & Paste HTML Code</h4><br>
                <form>
                    <div class="input-group">
                        <input type="text" class="form-control" value="" placeholder="code" id="snippet-html">
                        <span class="input-group-btn">
                            <button class="btn btn-default" type="button" id="copy-button-html" data-toggle="popover" data-placement="bottom" data-container="body" data-content="">Copy</button>
                        </span>
                    </div>
                </form><br><h4>Copy & Paste Markdown Code</h4><br>
                <form>
                    <div class="input-group">
                        <input type="text" class="form-control" value="" placeholder="code" id="snippet-md">
                        <span class="input-group-btn">
                            <button class="btn btn-default" type="button" id="copy-button-md" data-toggle="popover" data-placement="bottom" data-container="body" data-content="">Copy</button>
                        </span>
                    </div>
                </form><br><br>`)

        $('#copy-button-html').popover({
            container: 'body'
        })

        $('#copy-button-md').popover({
            container: 'body'
        })

        $('#copy-button-html').bind('click', function () {
            let input = document.getElementById('snippet-html');
            input.focus()
            input.setSelectionRange(0, input.value.length + 1)
            try {
                var success = document.execCommand('copy')
                if (success) {
                    $('#copy-button-html').trigger('copied', ['Copied!'])
                } else {
                    $('#copy-button-html').trigger('copied', ['Copy with Ctrl-c'])
                }
            } catch (err) {
                $('#copy-button-html').trigger('copied', ['Copy with Ctrl-c'])
            }
        });

        $('#copy-button-md').bind('click', function () {
            let input = document.getElementById('snippet-md');
            input.focus()
            input.setSelectionRange(0, input.value.length + 1)
            try {
                var success = document.execCommand('copy')
                if (success) {
                    $('#copy-button-md').trigger('copied', ['Copied!'])
                } else {
                    $('#copy-button-md').trigger('copied', ['Copy with Ctrl-c'])
                }
            } catch (err) {
                $('#copy-button-md').trigger('copied', ['Copy with Ctrl-c'])
            }
        });

        $('#copy-button-html').bind('copied', function (event, message) {
            $(this).attr('data-content', message)
                .popover('show')
                .attr('data-content', "")
        });

        $('#copy-button-md').bind('copied', function (event, message) {
            $(this).attr('data-content', message)
                .popover('show')
                .attr('data-content', "")
        });

    })
}

// getClassByContent returns the matching css class depending on a label's content
function getClassByContent(label) {
    if (label.includes("-")) {
        return 'badge badge-danger'
    } else if (label.includes("+")) {
        return 'badge badge-success'
    } else if (label.includes("!")) {
        return 'badge badge-warning'
    } else {
        return 'badge badge-secondary'
    }
}

// getCodeUrlencoded returns a url encoded label code
function getCodeUrlencoded() {
    let label = code_label.join('_')
    let encoded = encodeURIComponent(label);
    return encoded;
}