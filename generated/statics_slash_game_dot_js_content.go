package generated
const statics_slash_game_dot_js = "var borderRadius = 8;\x0Avar borderMargin = 4;\x0Avar borderWidth = 4;\x0Avar pieceMargin = 3;\x0Avar pieceWidth = 4;\x0Avar horizon = 4;\x0Avar vertial = 5;\x0Avar sizes = [[2, 2], [2, 1], [1, 2], [1, 1]];\x0Avar colors = ['rgb(128,0,0,', 'rgb(0,0,128,', 'rgb(0,128,0,', 'rgb(128,128,0,'];\x0Afunction color(i, alpha) {\x0A    return colors[i] + alpha + \x22)\x22;\x0A}\x0Avar pieceLineColor = 'rgb(0,0,0)';\x0Avar masks = ['2x2', '2x1', '1x2', '1x1'];\x0Avar Board = /** @class */ (function () {\x0A    function Board() {\x0A        this.reset();\x0A    }\x0A    Board.prototype.reset = function () {\x0A        this.used = new Array();\x0A        this.sizes = new Array();\x0A        this.pieces = new Array();\x0A        this.evts = new Array();\x0A        this.moving = false;\x0A    };\x0A    Board.prototype.piece = function (cell, p) {\x0A        this.pieces[cell] = p;\x0A    };\x0A    Board.prototype.getSize = function (cell) {\x0A        switch (this.sizes[cell]) {\x0A            case '2x2':\x0A                return [2, 2];\x0A            case '2x1':\x0A                return [2, 1];\x0A            case '1x2':\x0A                return [1, 2];\x0A            case '1x1':\x0A                return [1, 1];\x0A        }\x0A    };\x0A    Board.prototype.serialize = function () {\x0A        var _this = this;\x0A        var all = '[';\x0A        var c = 0;\x0A        masks.forEach(function (mask) {\x0A            if (c != 0) {\x0A                all = all + ',';\x0A            }\x0A            all = all + '[';\x0A            var counter = 0;\x0A            for (var k in _this.sizes) {\x0A                if (_this.sizes[k] == mask) {\x0A                    if (counter != 0) {\x0A                        all = all + ',';\x0A                    }\x0A                    all = all + k;\x0A                    counter++;\x0A                }\x0A            }\x0A            all = all + ']';\x0A            c++;\x0A        });\x0A        return all + ']';\x0A    };\x0A    Board.prototype.toIndex = function (h, v) {\x0A        return v * horizon + h;\x0A    };\x0A    Board.prototype.toXY = function (idx) {\x0A        return new XYPair(idx % horizon, Math.floor(idx / horizon));\x0A    };\x0A    Board.prototype.can = function (w, h, cell) {\x0A        if ((w == 2 && cell % horizon == horizon - 1) || (h == 2 && Math.floor(cell / horizon) == vertial - 1)) {\x0A            return false;\x0A        }\x0A        for (var j = 0; j < h; j++) {\x0A            for (var i = 0; i < w; i++) {\x0A                var t = cell + i;\x0A                if (t in this.used) {\x0A                    return false;\x0A                }\x0A            }\x0A            cell = cell + horizon;\x0A        }\x0A        return true;\x0A    };\x0A    Board.prototype.put = function (w, h, cell, f) {\x0A        var rawCell = cell;\x0A        for (var j = 0; j < h; j++) {\x0A            for (var i = 0; i < w; i++) {\x0A                this.used[cell + i] = rawCell;\x0A            }\x0A            cell = cell + horizon;\x0A        }\x0A        this.sizes[rawCell] = w + \x22x\x22 + h;\x0A        this.evts[rawCell] = f;\x0A    };\x0A    Board.prototype.remove = function (w, h, cell) {\x0A        var rawCell = cell;\x0A        for (var j = 0; j < h; j++) {\x0A            for (var i = 0; i < w; i++) {\x0A                var k = cell + i;\x0A                delete (this.used[k]);\x0A            }\x0A            cell = cell + horizon;\x0A        }\x0A        delete (this.sizes[rawCell]);\x0A        delete (this.pieces[rawCell]);\x0A        delete (this.evts[rawCell]);\x0A    };\x0A    Board.prototype.move = function (w, h, dst, src) {\x0A        var rawCell = src;\x0A        for (var j = 0; j < h; j++) {\x0A            for (var i = 0; i < w; i++) {\x0A                delete (this.used[src + i]);\x0A            }\x0A            src = src + horizon;\x0A        }\x0A        var p = this.pieces[rawCell];\x0A        delete (this.sizes[rawCell]);\x0A        delete (this.pieces[rawCell]);\x0A        var f = this.evts[rawCell];\x0A        delete (this.evts[rawCell]);\x0A        this.put(w, h, dst, f);\x0A        this.piece(dst, p);\x0A        f(dst);\x0A    };\x0A    Board.prototype.canMove = function (w, h, cell) {\x0A        var allMoves = new Array();\x0A        allMoves[cell] = 1;\x0A        var tested = new Array();\x0A        while (true) {\x0A            var newMoves = new Array();\x0A            for (var k in allMoves) {\x0A                if (k in tested) {\x0A                    continue;\x0A                }\x0A                var moves = this.canMoveImpl(w, h, parseInt(k));\x0A                for (var j in moves) {\x0A                    newMoves[moves[j]] = 1;\x0A                }\x0A                tested[k] = 1;\x0A            }\x0A            if (newMoves.length == 0) {\x0A                break;\x0A            }\x0A            for (var k in newMoves) {\x0A                allMoves[k] = 1;\x0A            }\x0A        }\x0A        var ret = new Array();\x0A        for (var k in allMoves) {\x0A            ret.push(parseInt(k));\x0A        }\x0A        return ret;\x0A    };\x0A    Board.prototype.canMoveImpl = function (w, h, cell) {\x0A        var ret = new Array();\x0A        ret.push(cell);\x0A        var up = true, down = true, left = true, right = true;\x0A        if (cell >= horizon) {\x0A            for (var i = 0; i < w; i++) {\x0A                var k = cell + i - horizon;\x0A                if ((k in this.used) && this.used[k] != cell) {\x0A                    up = false;\x0A                    break;\x0A                }\x0A            }\x0A        }\x0A        else {\x0A            up = false;\x0A        }\x0A        if (cell < horizon * (vertial - h)) {\x0A            for (var i = 0; i < w; i++) {\x0A                var k = cell + i + h * horizon;\x0A                if ((k in this.used) && this.used[k] != cell) {\x0A                    down = false;\x0A                    break;\x0A                }\x0A            }\x0A        }\x0A        else {\x0A            down = false;\x0A        }\x0A        if (0 != cell % horizon) {\x0A            for (var i = 0; i < h; i++) {\x0A                var k = cell + i * horizon - 1;\x0A                if ((k in this.used) && this.used[k] != cell) {\x0A                    left = false;\x0A                    break;\x0A                }\x0A            }\x0A        }\x0A        else {\x0A            left = false;\x0A        }\x0A        if (cell % horizon < horizon - w) {\x0A            for (var i = 0; i < h; i++) {\x0A                var k = cell + w + i * horizon;\x0A                if ((k in this.used) && this.used[k] != cell) {\x0A                    right = false;\x0A                    break;\x0A                }\x0A            }\x0A        }\x0A        else {\x0A            right = false;\x0A        }\x0A        if (up) {\x0A            ret.push(cell - horizon);\x0A        }\x0A        if (down) {\x0A            ret.push(cell + horizon);\x0A        }\x0A        if (left) {\x0A            ret.push(cell - 1);\x0A        }\x0A        if (right) {\x0A            ret.push(cell + 1);\x0A        }\x0A        return ret;\x0A    };\x0A    Board.prototype.setDragable = function (e) {\x0A        for (var k in this.pieces) {\x0A            this.pieces[k].setDragable(e);\x0A        }\x0A    };\x0A    Board.prototype.win = function (cell) {\x0A        return cell == 13 && this.sizes[cell] === '2x2';\x0A    };\x0A    return Board;\x0A}());\x0Avar CanvasContext = /** @class */ (function () {\x0A    function CanvasContext(canvas, ctx) {\x0A        this.canvas = canvas;\x0A        this.ctx = ctx;\x0A        canvasProperties.checkTransform(ctx);\x0A    }\x0A    CanvasContext.prototype.getTransform = function () {\x0A        if (this.ctx.getTransform) {\x0A            return this.ctx.getTransform();\x0A        }\x0A        return null;\x0A    };\x0A    return CanvasContext;\x0A}());\x0Afunction main(cvs) {\x0A    // testComponents(cvs)\x0A    // testSprite(cvs)\x0A    // return\x0A    var canvas = document.getElementById(cvs);\x0A    var ctx = canvas.getContext('2d');\x0A    ctx.font = '16px Arial';\x0A    var board = new Board();\x0A    var cc = new CanvasContext(canvas, ctx);\x0A    ui(cc, board);\x0A}\x0Afunction request(url, param, f) {\x0A    var http = new XMLHttpRequest();\x0A    http.open(\x22GET\x22, url + \x22?\x22 + encodeURIComponent(param));\x0A    http.onreadystatechange = function (e) {\x0A        if (http.readyState == 4) {\x0A            f(http.status == 200, http.responseText);\x0A        }\x0A    };\x0A    http.send();\x0A}\x0Avar enableAllButton;\x0Afunction ui(cc, game) {\x0A    var canvas = cc.canvas;\x0A    var ctx = cc.ctx;\x0A    var widthAny = canvas.width;\x0A    var height = canvas.height;\x0A    var bg = new CanvasSprite(cc.getTransform(), 1, function (c, r) {\x0A        ctx.fillStyle = 'rgb(127,127,127)';\x0A        ctx.fillRect(0, 0, canvas.width, canvas.height);\x0A    });\x0A    var width = height * horizon / vertial;\x0A    var cellWidth = (width - borderMargin * 2 - ctx.lineWidth - 2 * borderRadius) / horizon;\x0A    var cellHeight = (height - borderMargin * 2 - ctx.lineWidth - 2 * borderRadius) / vertial;\x0A    var gridPoints = calculateGrids(cellWidth, cellHeight);\x0A    var board = new CanvasSprite(cc.getTransform(), 2, function (c, r) {\x0A        drawBorder(c, width, height);\x0A        drawGrid(c, width, height, gridPoints);\x0A    });\x0A    bg.add(board);\x0A    var uiLeft, uiRight, uiBottom, uiTop;\x0A    var mng = new OperatableRectMng(canvas);\x0A    var tmplargin = (height - borderMargin * 2 - cellHeight * 3 - borderRadius * 2) / 3;\x0A    var tmpls = new Array();\x0A    for (var sizeIdxKey in sizes) {\x0A        var sizeIdx = parseInt(sizeIdxKey);\x0A        var drawTmpl = function (sizeIdx) {\x0A            var size = sizes[sizeIdx];\x0A            var tmplRect = pieceDragRect(size[0], size[1], cellWidth * horizon + borderMargin * 2 + borderRadius * 2 + tmplargin / 5, borderMargin + borderRadius + tmplargin / 5, cellWidth, cellHeight);\x0A            if (sizeIdx == 1 || sizeIdx == 3) {\x0A                var dy = cellHeight * 2 + tmplargin / 5;\x0A                tmplRect.top += dy;\x0A                tmplRect.bottom += dy;\x0A            }\x0A            if (sizeIdx == 2 || sizeIdx == 3) {\x0A                var dx = cellWidth * 2 + tmplargin / 5;\x0A                tmplRect.left += dx;\x0A                tmplRect.right += dx;\x0A            }\x0A            if (sizeIdx == 0) {\x0A                uiLeft = tmplRect.left;\x0A                uiTop = tmplRect.top;\x0A            }\x0A            else if (sizeIdx == 1) {\x0A                uiBottom = tmplRect.bottom;\x0A            }\x0A            else if (sizeIdx == 3) {\x0A                uiRight = tmplRect.right;\x0A            }\x0A            var tmplZ = 30;\x0A            var shadow = null;\x0A            var tmpl = new CanvasMovableSprite(mng, cc.getTransform(), tmplZ, tmplRect, function (c, r) {\x0A                if (shadow != null) {\x0A                    var xy = game.toXY(shadow[0]);\x0A                    drawPiece(size[0], size[1], c, gridPoints[0][xy.x] + pieceMargin + ctx.lineWidth / 2, gridPoints[1][xy.y] + pieceMargin + ctx.lineWidth / 2, pieceLineColor, color(sizeIdx, 0.5), cellWidth, cellHeight);\x0A                }\x0A                var newRect = tmpl.getRect();\x0A                drawPiece(size[0], size[1], c, newRect.left, newRect.top, pieceLineColor, color(sizeIdx, 1), cellWidth, cellHeight);\x0A            }, function (downX, downY, x, y) {\x0A                tmpl.setZ(tmplZ + 5);\x0A                var i = inGrid(size[0], size[1], tmpl.getRect(), game, gridPoints);\x0A                if (i != null && game.can(size[0], size[1], i[0])) {\x0A                    shadow = i;\x0A                }\x0A                else {\x0A                    shadow = null;\x0A                }\x0A            }, function () {\x0A                var i = inGrid(size[0], size[1], tmpl.getRect(), game, gridPoints);\x0A                if (i != null) {\x0A                    if (game.can(size[0], size[1], i[0])) {\x0A                        createPiece(mng, bg, cc, game, sizeIdx, size[0], size[1], i[0], gridPoints, cellWidth, cellHeight, tmpl, solveButton, canvas);\x0A                        if (sizeIdx == 0) {\x0A                            tmpl.setDragable(false);\x0A                            solveButton.Enable(true);\x0A                        }\x0A                    }\x0A                }\x0A                shadow = null;\x0A                tmpl.setRect(tmplRect.left, tmplRect.top, tmplRect.right, tmplRect.bottom);\x0A                tmpl.setZ(tmplZ);\x0A            }, function () {\x0A                shadow = null;\x0A                tmpl.setRect(tmplRect.left, tmplRect.top, tmplRect.right, tmplRect.bottom);\x0A                tmpl.setZ(tmplZ);\x0A            });\x0A            tmpl.getOperable().setRecentRectWhenRelease(false);\x0A            bg.add(tmpl);\x0A            tmpls[sizeIdx] = tmpl;\x0A        };\x0A        drawTmpl(sizeIdx);\x0A    }\x0A    var buttonWidth = (uiRight - uiLeft + borderRadius) / 2 - tmplargin / 5 / 2;\x0A    var buttonHeight = cellHeight / 2;\x0A    var buttonTop = uiBottom + borderRadius * 2;\x0A    var buttonMargin = tmplargin / 5;\x0A    enableAllButton = function (e) {\x0A        solveButton.Enable(e);\x0A        easyButton.Enable(e);\x0A        mediumButton.Enable(e);\x0A        hardButton.Enable(e);\x0A    };\x0A    var procNewBoard = function (level) {\x0A        for (var k in game.pieces) {\x0A            game.pieces[k].setDragable(false);\x0A        }\x0A        for (var k in tmpls) {\x0A            tmpls[k].setDragable(false);\x0A        }\x0A        enableAllButton(false);\x0A        request(level, \x22\x22, function (succ, resp) {\x0A            if (succ) {\x0A                var result;\x0A                try {\x0A                    result = JSON.parse(resp);\x0A                    for (var k in game.pieces) {\x0A                        bg.remove(game.pieces[k]);\x0A                    }\x0A                    game.reset();\x0A                    var sizeIdx = 0;\x0A                    for (var i in result) {\x0A                        for (var j in result[i]) {\x0A                            createPiece(mng, bg, cc, game, sizeIdx, sizes[sizeIdx][0], sizes[sizeIdx][1], result[i][j], gridPoints, cellWidth, cellHeight, tmpls[sizeIdx], solveButton, canvas);\x0A                        }\x0A                        sizeIdx++;\x0A                    }\x0A                }\x0A                catch (e) {\x0A                    alert(e);\x0A                }\x0A            }\x0A            else {\x0A                alert(resp);\x0A            }\x0A            enableAllButton(true);\x0A            for (var k in tmpls) {\x0A                tmpls[k].setDragable(true);\x0A            }\x0A        });\x0A    };\x0A    var easyButton = drawButton(cc, bg, mng, new CanvasRect(uiLeft, buttonTop, uiLeft + buttonWidth, buttonTop + buttonHeight), bg.getZ() + 10, \x22Random Easy\x22, function () {\x0A        procNewBoard(\x22easy\x22);\x0A    }, function () { });\x0A    easyButton.Enable(true);\x0A    var mediumButton = drawButton(cc, bg, mng, new CanvasRect(uiLeft + buttonWidth + buttonMargin, buttonTop, uiLeft + buttonWidth + buttonMargin + buttonWidth, buttonTop + buttonHeight), bg.getZ() + 10, \x22Random Medium\x22, function () {\x0A        procNewBoard(\x22medium\x22);\x0A    }, function () { });\x0A    mediumButton.Enable(true);\x0A    var hardButton = drawButton(cc, bg, mng, new CanvasRect(uiLeft, buttonTop + buttonHeight + buttonMargin, uiLeft + buttonWidth, buttonTop + buttonHeight + buttonMargin + buttonHeight), bg.getZ() + 10, \x22Random Hard\x22, function () {\x0A        procNewBoard(\x22hard\x22);\x0A    }, function () { });\x0A    hardButton.Enable(true);\x0A    var solveButton = drawButton(cc, bg, mng, new CanvasRect(uiLeft + buttonWidth + buttonMargin, buttonTop + buttonHeight + buttonMargin, uiLeft + buttonWidth + buttonMargin + buttonWidth, buttonTop + buttonHeight + buttonMargin + buttonHeight), bg.getZ() + 10, \x22Solve\x22, function () {\x0A        game.setDragable(false);\x0A        enableAllButton(false);\x0A        request(\x22solve\x22, game.serialize(), function (succ, resp) {\x0A            if (succ) {\x0A                var result = JSON.parse(resp);\x0A                playPaths(cc, bg, mng, game, gridPoints, result, uiLeft - borderRadius / 2, uiTop, uiRight + borderMargin + borderRadius, buttonTop + buttonHeight + buttonMargin + buttonHeight, buttonWidth, buttonHeight, buttonMargin, function () {\x0A                    game.setDragable(true);\x0A                    enableAllButton(true);\x0A                });\x0A            }\x0A            else {\x0A                alert(resp);\x0A                game.setDragable(true);\x0A                enableAllButton(true);\x0A            }\x0A        });\x0A    }, function () { });\x0A    solveButton.Enable(false);\x0A    RootSprite(bg, ctx);\x0A}\x0Afunction playPaths(cc, bg, mng, game, gridPoints, result, uiLeft, uiTop, uiRight, uiBottom, buttonWidth, buttonHeight, buttonMargin, exitCallback) {\x0A    var ctx = cc.ctx;\x0A    ctx.save();\x0A    var b = new Button(function () {\x0A        bg.remove(bk);\x0A        mng.remove(b);\x0A        game.setDragable(true);\x0A        exitCallback();\x0A    }, function () { });\x0A    b.Enable(true);\x0A    b.setRect(uiLeft, uiTop, uiRight, uiBottom);\x0A    b.setZ(10000);\x0A    b.setTransform(cc.getTransform());\x0A    // mng.add(b)\x0A    var text;\x0A    var bk = new CanvasButtonSprite(mng, cc.getTransform(), 10000, function () {\x0A        ctx.fillStyle = 'rgb(128,128,128,0.95)';\x0A        ctx.fillRect(uiLeft, uiTop, uiRight - uiLeft, uiBottom - uiTop);\x0A        if (text != \x22\x22) {\x0A            ctx.font = '36px Arial';\x0A            ctx.fillStyle = 'rgb(0,0,0)';\x0A            ctx.textAlign = 'center';\x0A            ctx.textBaseline = 'middle';\x0A            ctx.strokeStyle = 'rgb(255,255,255)';\x0A            ctx.fillText(text, prevButtonLeft + buttonWidth + buttonMargin / 2, 200);\x0A        }\x0A    }, b);\x0A    bg.add(bk);\x0A    var prevButtonLeft = uiLeft + borderMargin;\x0A    if (result.win) {\x0A        text = \x22Already solved\x22;\x0A    }\x0A    else {\x0A        if (result.path.length == 0) {\x0A            text = \x22No solution\x22;\x0A        }\x0A        else {\x0A            var curStep = 0;\x0A            var notMoving = true;\x0A            var display = function () {\x0A                text = curStep + \x22/\x22 + result.path.length + \x22 step\x22 + (result.path.length > 1 ? \x22s\x22 : \x22\x22);\x0A                prevButton.Enable(notMoving && curStep != 0);\x0A                nextButton.Enable(notMoving && curStep != result.path.length);\x0A            };\x0A            var movePiece = function (from, to) {\x0A                game.moving = true;\x0A                enableAllButton(false);\x0A                prevButton.Enable(false);\x0A                nextButton.Enable(false);\x0A                b.Enable(false);\x0A                var p = game.pieces[from];\x0A                var rawxy = game.toXY(from);\x0A                var toxy = game.toXY(to);\x0A                var moveStep = 5;\x0A                var dx = (gridPoints[0][toxy.x] - gridPoints[0][rawxy.x]) / moveStep;\x0A                var dy = (gridPoints[1][toxy.y] - gridPoints[1][rawxy.y]) / moveStep;\x0A                var counter = 0;\x0A                var period = 0.3 / moveStep;\x0A                var now = (new Date().getTime)();\x0A                var r = p.getRect().copy();\x0A                notMoving = false;\x0A                var timerHandle = setInterval(function () {\x0A                    if (!game.moving) {\x0A                        clearInterval(timerHandle);\x0A                        notMoving = true;\x0A                        return;\x0A                    }\x0A                    counter = ((new Date().getTime)() - now) / (period * 1000);\x0A                    if (counter >= moveStep) {\x0A                        counter = moveStep;\x0A                    }\x0A                    p.setRect(r.left + counter * dx, r.top + counter * dy, r.right + counter * dx, r.bottom + counter * dy);\x0A                    if (counter >= moveStep) {\x0A                        var wh = game.getSize(from);\x0A                        game.move(wh[0], wh[1], to, from);\x0A                        // prevButton.Enable(true)\x0A                        // nextButton.Enable(true)\x0A                        b.Enable(true);\x0A                        enableAllButton(true);\x0A                        clearInterval(timerHandle);\x0A                        game.moving = false;\x0A                        notMoving = true;\x0A                        display();\x0A                    }\x0A                }, period * 1000);\x0A            };\x0A            var prevButton = drawButton(cc, bk, mng, new CanvasRect(prevButtonLeft, 300, uiLeft + borderMargin + buttonWidth, 300 + buttonHeight), bk.getZ() + 10, \x22Prev\x22, function () {\x0A                if (curStep > 0) {\x0A                    var step = result.path[curStep - 1];\x0A                    movePiece(step[1], step[0]);\x0A                    curStep--;\x0A                }\x0A                display();\x0A            }, function () { });\x0A            var nextButton = drawButton(cc, bk, mng, new CanvasRect(uiLeft + buttonWidth + borderMargin + buttonMargin, 300, uiLeft + buttonWidth + borderMargin + buttonMargin + buttonWidth, 300 + buttonHeight), bk.getZ() + 10, \x22Next\x22, function () {\x0A                if (curStep != result.path.length) {\x0A                    var step = result.path[curStep];\x0A                    movePiece(step[0], step[1]);\x0A                    curStep++;\x0A                }\x0A                display();\x0A            }, function () { });\x0A            nextButton.Enable(true);\x0A            display();\x0A        }\x0A    }\x0A    ctx.restore();\x0A}\x0Afunction drawButton(cc, bg, mng, rect, z, text, clickFunc, cancelFunc) {\x0A    var ctx = cc.ctx;\x0A    var b = new Button(clickFunc, cancelFunc);\x0A    b.setRect(rect.left, rect.top, rect.right, rect.bottom);\x0A    b.setZ(z);\x0A    b.setTransform(cc.getTransform());\x0A    var s = new CanvasButtonSprite(mng, cc.getTransform(), z, function () {\x0A        ctx.fillStyle = 'rgb(0,0,0)';\x0A        ctx.textAlign = 'center';\x0A        ctx.textBaseline = 'middle';\x0A        ctx.fillRect(rect.left, rect.top, rect.width(), rect.height());\x0A        ctx.fillStyle = 'rgb(255,255,255)';\x0A        if (b.enable) {\x0A            ctx.fillText(text, rect.left + (rect.right - rect.left) / 2, rect.top + (rect.bottom - rect.top) / 2);\x0A        }\x0A        else {\x0A            ctx.strokeStyle = 'rgb(128,128,128)';\x0A            ctx.strokeText(text, rect.left + (rect.right - rect.left) / 2, rect.top + (rect.bottom - rect.top) / 2);\x0A        }\x0A    }, b);\x0A    bg.add(s);\x0A    // mng.add(b)\x0A    return b;\x0A}\x0Avar pieceZ = 20;\x0Afunction createPiece(mng, bg, cc, game, sizeIdx, w, h, cell, gridPoints, cellWidth, cellHeight, tmpl, solveButton, canvas) {\x0A    var ctx = cc.ctx;\x0A    var xy = game.toXY(cell);\x0A    var pieceRect = pieceDragRect(w, h, gridPoints[0][xy.x], gridPoints[1][xy.y], cellWidth, cellHeight);\x0A    var shadow = null;\x0A    var moves = null;\x0A    var piece = new CanvasMovableSprite(mng, cc.getTransform(), pieceZ, pieceRect, function (c, r) {\x0A        var newRect = piece.getRect();\x0A        if (shadow != null) {\x0A            var xy = game.toXY(shadow[0]);\x0A            if (moves != null) {\x0A                moves = game.canMove(w, h, cell);\x0A                if (-1 != moves.indexOf(shadow[0])) {\x0A                    drawPiece(w, h, c, gridPoints[0][xy.x] + pieceMargin + ctx.lineWidth / 2, gridPoints[1][xy.y] + pieceMargin + ctx.lineWidth / 2, pieceLineColor, color(sizeIdx, 0.5), cellWidth, cellHeight);\x0A                }\x0A            }\x0A        }\x0A        drawPiece(w, h, ctx, newRect.left, newRect.top, pieceLineColor, color(sizeIdx, 1), cellWidth, cellHeight);\x0A    }, function () {\x0A        piece.setZ(pieceZ + 100);\x0A        var newRect = piece.getRect();\x0A        var i = inGrid(w, h, newRect, game, gridPoints);\x0A        if (i != null) {\x0A            shadow = i;\x0A            if (moves == null) {\x0A                moves = game.canMove(w, h, cell);\x0A            }\x0A        }\x0A        else {\x0A            shadow = null;\x0A            moves = null;\x0A        }\x0A    }, function () {\x0A        var newRect = piece.getRect();\x0A        var i = inGrid(w, h, newRect, game, gridPoints);\x0A        if (i != null) {\x0A            if (moves == null) {\x0A                moves = game.canMove(w, h, cell);\x0A            }\x0A            if (-1 != moves.indexOf(i[0])) {\x0A                var xy = game.toXY(i[0]);\x0A                var newRectWidth = newRect.width();\x0A                var newRectHeight = newRect.height();\x0A                newRect.left = gridPoints[0][xy.x] + pieceMargin + ctx.lineWidth / 2;\x0A                newRect.top = gridPoints[1][xy.y] + pieceMargin + ctx.lineWidth / 2;\x0A                newRect.right = newRect.left + newRectWidth;\x0A                newRect.bottom = newRect.top + newRectHeight;\x0A                piece.setRect(newRect.left, newRect.top, newRect.right, newRect.bottom);\x0A                game.move(w, h, i[0], cell);\x0A                cell = i[0];\x0A                if (game.win(cell)) {\x0A                    createWinPanel(mng, bg, cc, game, canvas);\x0A                }\x0A            }\x0A            else {\x0A                newRect = pieceRect.copy();\x0A            }\x0A        }\x0A        else {\x0A            bg.remove(piece);\x0A            game.remove(w, h, cell);\x0A            if (sizeIdx == 0) {\x0A                tmpl.setDragable(true);\x0A                solveButton.Enable(false);\x0A            }\x0A        }\x0A        pieceRect = newRect.copy();\x0A        shadow = null;\x0A        piece.setRect(newRect.left, newRect.top, newRect.right, newRect.bottom);\x0A        piece.setZ(pieceZ);\x0A        moves = null;\x0A    }, function () {\x0A        shadow = null;\x0A        piece.setRect(pieceRect.left, pieceRect.top, pieceRect.right, pieceRect.bottom);\x0A        piece.setZ(pieceZ);\x0A        moves = null;\x0A    });\x0A    game.put(w, h, cell, function (pos) {\x0A        cell = pos;\x0A    });\x0A    bg.add(piece);\x0A    game.piece(cell, piece);\x0A}\x0Afunction createWinPanel(mng, bg, cc, game, canvas) {\x0A    var ctx = cc.ctx;\x0A    var uiLeft = 0, uiTop = 0, uiRight = canvas.width, uiBottom = canvas.height;\x0A    var text = \x22You win\x22;\x0A    var b = new Button(function () {\x0A        bg.remove(bk);\x0A        mng.remove(b);\x0A        game.setDragable(true);\x0A    }, function () { });\x0A    b.Enable(true);\x0A    b.setRect(uiLeft, uiTop, uiRight, uiBottom);\x0A    b.setTransform(cc.getTransform());\x0A    var bk = new CanvasButtonSprite(mng, cc.getTransform(), 10000, function () {\x0A        ctx.fillStyle = 'rgb(128,128,128,0.95)';\x0A        ctx.fillRect(uiLeft, uiTop, uiRight - uiLeft, uiBottom - uiTop);\x0A        if (text) {\x0A            ctx.font = '36px Arial';\x0A            ctx.fillStyle = 'rgb(0,0,0)';\x0A            ctx.textAlign = 'center';\x0A            ctx.textBaseline = 'middle';\x0A            ctx.strokeStyle = 'rgb(255,255,255)';\x0A            ctx.fillText(text, uiLeft + (uiRight - uiLeft) / 2, 200);\x0A        }\x0A    }, b);\x0A    bg.add(bk);\x0A    b.setZ(bk.getZ());\x0A    // mng.add(b)\x0A}\x0Afunction inGrid(w, h, r, game, gridPoints) {\x0A    var l = gridPoints[0].length;\x0A    var horIn = false, verIn = false;\x0A    for (var i = 0; i < l - w; i++) {\x0A        if (gridPoints[0][i] < r.left && r.left < gridPoints[0][i + 1]) {\x0A            horIn = i;\x0A            if (i != l - w - 1 && Math.abs(gridPoints[0][i + 1] - r.left) < Math.abs(gridPoints[0][i] - r.left)) {\x0A                horIn = i + 1;\x0A            }\x0A            break;\x0A        }\x0A    }\x0A    if (horIn === false && (gridPoints[0][w - 1] < r.right && r.right < gridPoints[0][w])) {\x0A        horIn = 0;\x0A    }\x0A    l = gridPoints[1].length;\x0A    for (var i = 0; i < l - h; i++) {\x0A        if (gridPoints[1][i] < r.top && r.top < gridPoints[1][i + 1]) {\x0A            verIn = i;\x0A            if (i != l - h - 1 && Math.abs(gridPoints[1][i + 1] - r.top) < Math.abs(gridPoints[1][i] - r.top)) {\x0A                verIn = i + 1;\x0A            }\x0A            break;\x0A        }\x0A    }\x0A    if (verIn === false && (gridPoints[1][h - 1] < r.bottom && r.bottom < gridPoints[1][h])) {\x0A        verIn = 0;\x0A    }\x0A    if (horIn !== false && verIn !== false) {\x0A        var ret = new Array();\x0A        for (var i = 0; i < w; i++) {\x0A            for (var j = 0; j < h; j++) {\x0A                var v = game.toIndex(i + horIn, j + verIn);\x0A                ret.push(v);\x0A            }\x0A        }\x0A        return ret;\x0A    }\x0A    else {\x0A        return null;\x0A    }\x0A}\x0Afunction drawBorder(ctx, width, height) {\x0A    ctx.save();\x0A    ctx.fillStyle = 'rgb(255,0,0,1)';\x0A    ctx.strokeStyle = 'rgb(0, 0, 0, 1)';\x0A    ctx.lineWidth = 4;\x0A    ctx.beginPath();\x0A    ctx.arc(borderRadius + borderMargin + ctx.lineWidth / 2, borderRadius + borderMargin + ctx.lineWidth / 2, borderRadius, Math.PI * 1.5, Math.PI, true);\x0A    ctx.lineTo(borderMargin + ctx.lineWidth / 2, height - borderRadius - borderMargin - ctx.lineWidth / 2);\x0A    ctx.arc(borderRadius + borderMargin + ctx.lineWidth / 2, height - borderRadius - borderMargin - ctx.lineWidth / 2, borderRadius, Math.PI, Math.PI / 2, true);\x0A    ctx.lineTo(width - 2 * borderMargin - ctx.lineWidth / 2 * 2 - 2 * borderRadius, height - borderMargin - ctx.lineWidth / 2);\x0A    ctx.arc(width - borderMargin - ctx.lineWidth / 2 - borderRadius, height - borderMargin - ctx.lineWidth / 2 - borderRadius, borderRadius, Math.PI / 2, 0, true);\x0A    ctx.lineTo(width - borderMargin - ctx.lineWidth / 2, borderMargin + ctx.lineWidth / 2 + borderRadius);\x0A    ctx.arc(width - borderMargin - ctx.lineWidth / 2 - borderRadius, borderMargin + ctx.lineWidth / 2 + borderRadius, borderRadius, 0, Math.PI * 1.5, true);\x0A    ctx.lineTo(borderMargin + ctx.lineWidth / 2 + borderRadius, borderMargin + ctx.lineWidth / 2);\x0A    ctx.stroke();\x0A    ctx.fill();\x0A    ctx.restore();\x0A}\x0Avar dashLineWidth = 2;\x0Afunction calculateGrids(cellWidth, cellHeight) {\x0A    var points = new Array(2);\x0A    points[0] = new Array(horizon);\x0A    points[1] = new Array(vertial);\x0A    for (var i = 0; i <= horizon; i++) {\x0A        points[0][i] = i * cellWidth + borderRadius + borderMargin + dashLineWidth / 2;\x0A    }\x0A    for (var i = 0; i <= vertial; i++) {\x0A        points[1][i] = i * cellHeight + borderRadius + borderMargin + dashLineWidth / 2;\x0A    }\x0A    return points;\x0A}\x0Afunction drawGrid(ctx, width, height, gridPoints) {\x0A    ctx.save();\x0A    ctx.strokeStyle = 'rgb(127, 127, 127, 0.5)';\x0A    ctx.lineWidth = dashLineWidth;\x0A    ctx.setLineDash([4, 2]);\x0A    ctx.beginPath();\x0A    for (var i = 0; i <= horizon; i++) {\x0A        var x = gridPoints[0][i];\x0A        ctx.moveTo(x, borderRadius + borderMargin + ctx.lineWidth / 2);\x0A        ctx.lineTo(x, height - borderMargin - borderRadius - ctx.lineWidth / 2);\x0A    }\x0A    for (var i = 0; i <= vertial; i++) {\x0A        var y = gridPoints[1][i];\x0A        ctx.moveTo(borderRadius + borderMargin + ctx.lineWidth / 2, y);\x0A        ctx.lineTo(width - borderRadius - borderMargin - ctx.lineWidth / 2, y);\x0A    }\x0A    ctx.stroke();\x0A    ctx.restore();\x0A}\x0Afunction pieceDragRect(w, h, x, y, cellWidth, cellHeight) {\x0A    return new CanvasRect(x + pieceMargin, y + pieceMargin, x + w * cellWidth - 2 * pieceMargin - pieceWidth, y + h * cellHeight - 2 * pieceMargin - pieceWidth);\x0A}\x0Afunction drawPiece(w, h, ctx, x, y, lineColor, fillColor, cellWidth, cellHeight) {\x0A    ctx.lineWidth = pieceWidth;\x0A    ctx.strokeStyle = lineColor;\x0A    ctx.fillStyle = fillColor;\x0A    ctx.beginPath();\x0A    ctx.arc(x + ctx.lineWidth / 2 + borderRadius, y + ctx.lineWidth / 2 + borderRadius, borderRadius, Math.PI * 1.5, Math.PI, true);\x0A    ctx.lineTo(x + ctx.lineWidth / 2, y + h * cellHeight - 2 * pieceMargin - ctx.lineWidth / 2 - borderRadius);\x0A    ctx.arc(x + ctx.lineWidth / 2 + borderRadius, y + h * cellHeight - 2 * pieceMargin - ctx.lineWidth / 2 - borderRadius, borderRadius, Math.PI, Math.PI / 2, true);\x0A    ctx.lineTo(x + w * cellWidth - 2 * pieceMargin - ctx.lineWidth / 2 - borderRadius, y + h * cellHeight - 2 * pieceMargin - ctx.lineWidth / 2);\x0A    ctx.arc(x + w * cellWidth - 2 * pieceMargin - ctx.lineWidth / 2 - borderRadius, y + h * cellHeight - 2 * pieceMargin - ctx.lineWidth / 2 - borderRadius, borderRadius, Math.PI / 2, 0, true);\x0A    ctx.lineTo(x + w * cellWidth - 2 * pieceMargin - ctx.lineWidth / 2, y + ctx.lineWidth / 2 + borderRadius);\x0A    ctx.arc(x + w * cellWidth - 2 * pieceMargin - ctx.lineWidth / 2 - borderRadius, y + ctx.lineWidth / 2 + borderRadius, borderRadius, 0, Math.PI * 1.5, true);\x0A    ctx.lineTo(x + ctx.lineWidth / 2 + borderRadius, y + ctx.lineWidth / 2);\x0A    ctx.fill();\x0A    ctx.stroke();\x0A}\x0A"
