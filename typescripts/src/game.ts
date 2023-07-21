const borderRadius = 8
const borderMargin=4
const borderWidth = 4

const pieceMargin = 3
const pieceWidth = 4

const horizon=4
const vertial=5

const sizes=[[2, 2], [2, 1], [1, 2], [1, 1]]
const colors=['rgb(128,0,0,', 'rgb(0,0,128,', 'rgb(0,128,0,', 'rgb(128,128,0,']
function color(i:number, alpha:number):string{
    return colors[i]+alpha+")"
}
const pieceLineColor='rgb(0,0,0)'

const masks=['2x2', '2x1', '1x2', '1x1']
class Board{
    moving:boolean
    used:Array<number>
    sizes:Array<string>
    pieces:Array<CanvasOperatableSprite>
    evts:Array<(number)=>void>
    constructor(){
        this.reset()
    }
    reset():void{
        this.used = new Array()
        this.sizes = new Array()
        this.pieces = new Array()
        this.evts = new Array()
        this.moving=false
    }
    piece(cell:number, p:CanvasOperatableSprite):void{
        this.pieces[cell]=p
    }
    getSize(cell:number):number[]{
        switch( this.sizes[cell] ){
            case '2x2':
                return [2,2]
            case '2x1':
                return [2,1]
            case '1x2':
                return [1,2]
            case '1x1':
                return [1,1]
        }
    }
    serialize():string{
        var all = '['
        var c=0
        masks.forEach(mask=>{
            if(c!=0){
                all = all+','
            }
            all = all+'['
            var counter=0
            for(var k in this.sizes){
                if(this.sizes[k]==mask){
                    if(counter!=0){
                        all = all + ','
                    }
                    all = all + k
                    counter++
                }
            }
            all = all + ']'
            c++
        })
        return all+']'
    }
    toIndex(h:number,v:number):number{
        return v*horizon+h
    }
    toXY(idx:number):XYPair{
        return new XYPair(idx%horizon, Math.floor(idx/horizon))
    }
    can(w:number, h:number, cell:number):boolean{
        if((w==2 && cell%horizon==horizon-1) || (h==2 && Math.floor(cell/horizon)==vertial-1)){
            return false
        }
        for(var j=0; j<h; j++){
            for(var i=0; i<w; i++){
                var t=cell+i
                if(t in this.used){
                    return false
                }
            }
            cell=cell+horizon
        }
        return true
    }
    put(w:number, h:number, cell:number, f:(number)=>void):void{
        var rawCell=cell
        for(var j=0; j<h; j++){
            for(var i=0; i<w; i++){
                this.used[cell+i]=rawCell
            }
            cell=cell+horizon
        }
        this.sizes[rawCell]=w+"x"+h
        this.evts[rawCell]=f
    }
    remove(w:number, h:number, cell:number):void{
        var rawCell=cell
        for(var j=0; j<h; j++){
            for(var i=0; i<w; i++){
                var k=cell+i
                delete(this.used[k])
            }
            cell=cell+horizon
        }
        delete(this.sizes[rawCell])
        delete(this.pieces[rawCell])
        delete(this.evts[rawCell])
    }
    move(w:number, h:number, dst:number, src:number):void{
        var rawCell=src
        for(var j=0; j<h; j++){
            for(var i=0; i<w; i++){
                delete(this.used[src+i])
            }
            src=src+horizon
        }
        var p=this.pieces[rawCell]
        delete(this.sizes[rawCell])
        delete(this.pieces[rawCell])
        var f=this.evts[rawCell]
        delete(this.evts[rawCell])
        this.put(w,h,dst,f)
        this.piece(dst, p)
        f(dst)
    }
    canMove(w:number, h:number, cell:number):Array<number>{
        var allMoves = new Array()
        allMoves[cell]=1
        var tested = new Array()
        while(true){
            var newMoves = new Array()
            for(var k in allMoves){
                if(k in tested){
                    continue
                }
                var moves=this.canMoveImpl(w,h,parseInt(k))
                for(var j in moves){
                    newMoves[moves[j]]=1
                }
                tested[k]=1
            }
            if(newMoves.length==0){
                break
            }
            for(var k in newMoves){
                allMoves[k]=1
            }
        }
        var ret=new Array<number>()
        for(var k in allMoves){
            ret.push(parseInt(k))
        }
        return ret
    }
    canMoveImpl(w:number, h:number, cell:number):Array<number>{
        var ret = new Array<number>()
        ret.push(cell)
        var up=true,down=true,left=true,right=true
        if(cell>=horizon){
            for(var i=0;i<w;i++){
                var k=cell+i-horizon
                if((k in this.used) && this.used[k]!=cell){
                    up=false
                    break
                }
            }
        }else{
            up=false
        }
        if(cell<horizon*(vertial-h)){
            for(var i=0;i<w;i++){
                var k=cell+i+h*horizon
                if((k in this.used) && this.used[k]!=cell){
                    down=false
                    break
                }
            }
        }else{
            down=false
        }
        if(0!=cell%horizon){
            for(var i=0;i<h;i++){
                var k=cell+i*horizon-1
                if((k in this.used) && this.used[k]!=cell){
                    left=false
                    break
                }
            }
        }else{
            left=false
        }
        if(cell%horizon < horizon-w){
            for(var i=0;i<h;i++){
                var k=cell+w+i*horizon
                if((k in this.used) && this.used[k]!=cell){
                    right=false
                    break
                }
            }
        }else{
            right=false
        }
        if(up){
            ret.push(cell-horizon)
        }
        if(down){
            ret.push(cell+horizon)
        }
        if(left){
            ret.push(cell-1)
        }
        if(right){
            ret.push(cell+1)
        }
        return ret
    }
    setDragable(e:boolean):void{
        for(var k in this.pieces){
            this.pieces[k].setDragable(e)
        }
    }
    win(cell:number):boolean{
        return cell==13 && this.sizes[cell]==='2x2'
    }
}
class CanvasContext{
    canvas:any
    ctx:any
    constructor(canvas:any, ctx:any){
        this.canvas=canvas
        this.ctx=ctx
        canvasProperties.checkTransform(ctx)
    }
    getTransform():any{
        if(this.ctx.getTransform){
            return this.ctx.getTransform()
        }
        return null
    }
}
function main(cvs:any):void{
    // testComponents(cvs)
    // testSprite(cvs)
    // return
    var canvas:any = document.getElementById(cvs);
    var ctx = canvas.getContext('2d')
    ctx.font = '16px Arial';
    var board=new Board()
    var cc=new CanvasContext(canvas, ctx)
    ui(cc, board)
}
function request(url:string, param:string, f:(boolean, string)=>void):void{
    const http = new XMLHttpRequest()
    http.open("GET", url+"?"+encodeURIComponent(param))
    http.onreadystatechange = function(e:any){
        if(http.readyState==4){
            f(http.status==200, http.responseText)
        }
    }
    http.send()
}
var enableAllButton:(boolean)=>void
function ui(cc:CanvasContext, game:Board):void{
    var canvas:any=cc.canvas
    var ctx:any=cc.ctx
    var widthAny=canvas.width
    var height=canvas.height
    var bg=new CanvasSprite(cc.getTransform(), 1, function(c, r){
        ctx.fillStyle='rgb(127,127,127)'
        ctx.fillRect(0, 0, canvas.width, canvas.height)
    })

    var width=(height as number)*horizon/vertial
    var cellWidth = (width-borderMargin*2-ctx.lineWidth-2*borderRadius)/horizon
    var cellHeight=((height as number)-borderMargin*2-ctx.lineWidth-2*borderRadius)/vertial
    var gridPoints = calculateGrids(cellWidth, cellHeight)
    var board=new CanvasSprite(cc.getTransform(), 2, function(c, r){
        drawBorder(c, width, height as number);
        drawGrid(c, width, height, gridPoints)
    })
    bg.add(board)

    var uiLeft:number, uiRight:number, uiBottom:number, uiTop:number
    var mng=new OperatableRectMng(canvas)
    var tmplargin = (height-borderMargin*2-cellHeight*3-borderRadius*2)/3
    var tmpls=new Array<CanvasMovableSprite>()
    for(var sizeIdxKey in sizes){
        var sizeIdx=parseInt(sizeIdxKey)
        var drawTmpl=function(sizeIdx:number):void{
            var size=sizes[sizeIdx]
            var tmplRect = pieceDragRect(size[0],size[1], cellWidth*horizon+borderMargin*2+borderRadius*2+tmplargin/5, borderMargin+borderRadius+tmplargin/5, cellWidth, cellHeight)
            if(sizeIdx==1 || sizeIdx==3){
                var dy=cellHeight*2+tmplargin/5
                tmplRect.top+=dy
                tmplRect.bottom+=dy
            }
            if(sizeIdx==2 || sizeIdx==3){
                var dx=cellWidth*2+tmplargin/5
                tmplRect.left+=dx
                tmplRect.right+=dx
            }
            if(sizeIdx==0){
                uiLeft=tmplRect.left
                uiTop=tmplRect.top
            }else if(sizeIdx==1){
                uiBottom=tmplRect.bottom
            }else if(sizeIdx==3){
                uiRight=tmplRect.right
            }
            const tmplZ=30
            var shadow:Array<number>=null
            var tmpl = new CanvasMovableSprite(mng, cc.getTransform(), tmplZ, tmplRect, 
                function(c:any, r:any):void{
                    if(shadow!=null){
                        var xy=game.toXY(shadow[0])
                        drawPiece(size[0], size[1], c, gridPoints[0][xy.x]+pieceMargin+ctx.lineWidth/2, gridPoints[1][xy.y]+pieceMargin+ctx.lineWidth/2, pieceLineColor, color(sizeIdx,0.5), cellWidth, cellHeight)
                    }
                    var newRect = tmpl.getRect()
                    drawPiece(size[0], size[1], c, newRect.left, newRect.top, pieceLineColor, color(sizeIdx,1), cellWidth, cellHeight)
                },
                function(downX:number, downY:number, x:number, y:number):void{
                    tmpl.setZ(tmplZ+5)
                    var i=inGrid(size[0], size[1], tmpl.getRect(), game, gridPoints)
                    if(i!=null && game.can(size[0], size[1], i[0])){
                        shadow=i
                    }else{
                        shadow=null
                    }
                },
                function():void{
                    var i=inGrid(size[0], size[1], tmpl.getRect(), game, gridPoints)
                    if(i!=null){
                        if(game.can(size[0], size[1], i[0])){
                            createPiece(mng, bg, cc, game, sizeIdx, size[0], size[1], i[0], gridPoints, cellWidth, cellHeight, tmpl, solveButton, canvas)
                            if(sizeIdx==0){
                                tmpl.setDragable(false)
                                solveButton.Enable(true)
                            }        
                        }
                    }
                    shadow=null
                    tmpl.setRect(tmplRect.left, tmplRect.top, tmplRect.right, tmplRect.bottom)
                    tmpl.setZ(tmplZ)
                },
                function():void{
                    shadow=null
                    tmpl.setRect(tmplRect.left, tmplRect.top, tmplRect.right, tmplRect.bottom)
                    tmpl.setZ(tmplZ)
                }
            )
            tmpl.getOperable().setRecentRectWhenRelease(false)
            bg.add(tmpl)
            tmpls[sizeIdx]=tmpl
        }
        drawTmpl(sizeIdx)
    }
    const buttonWidth=(uiRight-uiLeft+borderRadius)/2-tmplargin/5/2
    const buttonHeight=cellHeight/2
    const buttonTop=uiBottom+borderRadius*2
    const buttonMargin=tmplargin/5
    enableAllButton=function(e:boolean):void{
        solveButton.Enable(e)
        easyButton.Enable(e)
        mediumButton.Enable(e)
        hardButton.Enable(e)
    }
    var procNewBoard = function(level):void{
        for(var k in game.pieces){
            game.pieces[k].setDragable(false)
        }
        for(var k in tmpls){
            tmpls[k].setDragable(false)
        }
        enableAllButton(false)
        request(level, "",function(succ:boolean, resp:string):void{
            if(succ){
                var result:any
                try{
                    result=JSON.parse(resp)
                    for(var k in game.pieces){
                        bg.remove(game.pieces[k])
                    }
                    game.reset()
                    var sizeIdx=0
                    for(var i in result){
                        for(var j in result[i]){
                            createPiece(mng, bg, cc, game, sizeIdx, sizes[sizeIdx][0], sizes[sizeIdx][1], result[i][j], gridPoints, cellWidth, cellHeight, tmpls[sizeIdx], solveButton, canvas)
                        }
                        sizeIdx++
                    }
                }catch(e:any){
                    alert(e)
                }
            }else{
                alert(resp)
            }
            enableAllButton(true)
            for(var k in tmpls){
                tmpls[k].setDragable(true)
            }
        })
    }
    var easyButton=drawButton(cc, bg, mng, new CanvasRect(uiLeft, buttonTop, uiLeft+buttonWidth, buttonTop+buttonHeight), bg.getZ()+10, "Random Easy", 
        function():void{
            procNewBoard("easy")
        },
        function():void{}
    )
    easyButton.Enable(true)
    var mediumButton=drawButton(cc, bg, mng, new CanvasRect(uiLeft+buttonWidth+buttonMargin, buttonTop, uiLeft+buttonWidth+buttonMargin+buttonWidth, buttonTop+buttonHeight), bg.getZ()+10, "Random Medium", 
        function():void{
            procNewBoard("medium")
        },
        function():void{}
    )
    mediumButton.Enable(true)
    var hardButton=drawButton(cc, bg, mng, new CanvasRect(uiLeft, buttonTop+buttonHeight+buttonMargin, uiLeft+buttonWidth, buttonTop+buttonHeight+buttonMargin+buttonHeight), bg.getZ()+10, "Random Hard", 
        function():void{
            procNewBoard("hard")
        },
        function():void{}
    )
    hardButton.Enable(true)
    var solveButton=drawButton(cc, bg, mng, new CanvasRect(uiLeft+buttonWidth+buttonMargin, buttonTop+buttonHeight+buttonMargin, uiLeft+buttonWidth+buttonMargin+buttonWidth, buttonTop+buttonHeight+buttonMargin+buttonHeight), bg.getZ()+10, "Solve", 
        function():void{
            game.setDragable(false)
            enableAllButton(false)
            request("solve", game.serialize(), function(succ:boolean, resp:string){
                if(succ){
                    var result=JSON.parse(resp)
                    playPaths(cc, bg, mng, game, gridPoints, result, uiLeft-borderRadius/2, uiTop, uiRight+borderMargin+borderRadius, buttonTop+buttonHeight+buttonMargin+buttonHeight, buttonWidth, buttonHeight, buttonMargin, function(){
                        game.setDragable(true)
                        enableAllButton(true)    
                    })
                }else{
                    alert(resp)
                    game.setDragable(true)
                    enableAllButton(true)
                }
            })
        },
        function():void{}
    )
    solveButton.Enable(false)
    RootSprite(bg, ctx)
}
function playPaths(cc:CanvasContext, bg:CanvasSprite, mng:OperatableRectMng, game:Board, gridPoints:Array<Array<number>>, result:any, uiLeft:number, uiTop:number, uiRight:number, uiBottom:number, buttonWidth:number, buttonHeight:number, buttonMargin:number, exitCallback:()=>void):void{
    var ctx=cc.ctx
    ctx.save()
    var b=new Button(function():void{
        bg.remove(bk)
        mng.remove(b)
        game.setDragable(true)
        exitCallback()
    }, function():void{})
    b.Enable(true)
    b.setRect(uiLeft, uiTop, uiRight, uiBottom)
    b.setZ(10000)
    b.setTransform(cc.getTransform())
    // mng.add(b)

    var text:string
    var bk=new CanvasButtonSprite(mng, cc.getTransform(), 10000, function(){
        ctx.fillStyle='rgb(128,128,128,0.95)'
        ctx.fillRect(uiLeft,uiTop, uiRight-uiLeft, uiBottom-uiTop)

        if(text != ""){
            ctx.font = '36px Arial';
            ctx.fillStyle='rgb(0,0,0)'
            ctx.textAlign='center'
            ctx.textBaseline = 'middle'
            ctx.strokeStyle='rgb(255,255,255)'
            ctx.fillText(text, prevButtonLeft+buttonWidth+buttonMargin/2, 200)
        }
    },b)
    bg.add(bk)

    const prevButtonLeft=uiLeft+borderMargin
    if(result.win){
        text="Already solved"
    }else{
        if(result.path.length==0){
            text="No solution"
        }else{
            var curStep=0
            var notMoving:boolean = true
            var display = function():void{
                text = curStep+"/"+result.path.length + " step" + (result.path.length>1 ? "s" : "")
                prevButton.Enable(notMoving && curStep!=0)
                nextButton.Enable(notMoving && curStep!=result.path.length)
            }
            var movePiece = function(from:number, to:number):void{
                game.moving=true
                enableAllButton(false)
                prevButton.Enable(false)
                nextButton.Enable(false)
                b.Enable(false)
                var p=game.pieces[from]
                var rawxy=game.toXY(from)
                var toxy=game.toXY(to)
                const moveStep=5
                var dx=(gridPoints[0][toxy.x]-gridPoints[0][rawxy.x])/moveStep
                var dy=(gridPoints[1][toxy.y]-gridPoints[1][rawxy.y])/moveStep
                var counter=0
                const period:number=0.3/moveStep
                var now = (new Date().getTime)()
                var r=p.getRect().copy()
                notMoving=false
                var timerHandle=setInterval(function():void{
                    if(!game.moving){
                        clearInterval(timerHandle)
                        notMoving=true
                        return
                    }
                    counter=((new Date().getTime)()-now)/(period*1000)
                    if(counter>=moveStep){
                        counter=moveStep
                    }
                    p.setRect(r.left+counter*dx, r.top+counter*dy, r.right+counter*dx, r.bottom+counter*dy)
                    if(counter>=moveStep){
                        var wh=game.getSize(from)
                        game.move(wh[0], wh[1], to, from)
                        // prevButton.Enable(true)
                        // nextButton.Enable(true)
                        b.Enable(true)        
                        enableAllButton(true)
                        clearInterval(timerHandle)
                        game.moving=false
                        notMoving=true
                        display();
                    }
                }, period*1000)
            }
            var prevButton=drawButton(cc, bk, mng, new CanvasRect(prevButtonLeft, 300, uiLeft+borderMargin+buttonWidth, 300+buttonHeight), bk.getZ()+10, "Prev",
                function(){
                    if( curStep>0){
                        var step=result.path[curStep-1]
                        movePiece(step[1], step[0])
                        curStep--
                    }
                    display()
                },
                function(){}
            )
            var nextButton=drawButton(cc, bk, mng, new CanvasRect(uiLeft+buttonWidth+borderMargin+buttonMargin, 300, uiLeft+buttonWidth+borderMargin+buttonMargin+buttonWidth, 300+buttonHeight), bk.getZ()+10, "Next",
                function():void{
                    if(curStep!=result.path.length){
                        var step=result.path[curStep]
                        movePiece(step[0], step[1])
                        curStep++
                    }
                    display()
                },
                function():void{}
            )
            nextButton.Enable(true)
            display()
        }
    }
    ctx.restore()
}
function drawButton(cc:CanvasContext, bg:CanvasSprite, mng:OperatableRectMng, rect:CanvasRect, z:number, text:string, clickFunc:() => void, cancelFunc:() => void):Button{
    var ctx:any=cc.ctx
    var b=new Button(clickFunc, cancelFunc)
    b.setRect(rect.left, rect.top, rect.right, rect.bottom)
    b.setZ(z)
    b.setTransform(cc.getTransform())
    var s=new CanvasButtonSprite(mng, cc.getTransform(), z, function(){
        ctx.fillStyle='rgb(0,0,0)'
        ctx.textAlign='center'
        ctx.textBaseline = 'middle'
        ctx.fillRect(rect.left, rect.top, rect.width(), rect.height())
        ctx.fillStyle='rgb(255,255,255)'
        if(b.enable){
            ctx.fillText(text, rect.left+(rect.right-rect.left)/2, rect.top+(rect.bottom-rect.top)/2)
        }else{
            ctx.strokeStyle='rgb(128,128,128)'
            ctx.strokeText(text, rect.left+(rect.right-rect.left)/2, rect.top+(rect.bottom-rect.top)/2)
        }
    }, b)
    bg.add(s)
    // mng.add(b)
    return b
}
const pieceZ=20
function createPiece(mng:OperatableRectMng, bg:CanvasSprite, cc:CanvasContext, game:Board, sizeIdx:number, w:number, h:number, cell:number, gridPoints:Array<Array<number>>, cellWidth:number, cellHeight:number, tmpl:CanvasMovableSprite, solveButton:Button, canvas:any){
    var ctx:any=cc.ctx
    var xy=game.toXY(cell)
    var pieceRect = pieceDragRect(w, h, gridPoints[0][xy.x], gridPoints[1][xy.y], cellWidth, cellHeight)
    var shadow:Array<number>=null
    var moves=null
    var piece = new CanvasMovableSprite(mng, cc.getTransform(), pieceZ, pieceRect,
        function(c:any, r:CanvasRect):void{
            var newRect = piece.getRect()
            if(shadow != null){
                var xy=game.toXY(shadow[0])
                if(moves!=null){
                    moves=game.canMove(w, h, cell)
                    if(-1 !=moves.indexOf(shadow[0])){
                        drawPiece(w, h, c, gridPoints[0][xy.x]+pieceMargin+ctx.lineWidth/2, gridPoints[1][xy.y]+pieceMargin+ctx.lineWidth/2, pieceLineColor, color(sizeIdx,0.5), cellWidth, cellHeight)
                    }
                }
            }
            drawPiece(w, h, ctx, newRect.left, newRect.top, pieceLineColor, color(sizeIdx, 1), cellWidth, cellHeight)
        },
        function():void{
            piece.setZ(pieceZ+100)
            var newRect = piece.getRect()
            var i=inGrid(w, h, newRect, game, gridPoints)
            if(i!=null){
                shadow=i
                if(moves==null){
                    moves=game.canMove(w, h, cell)
                }
            }else{
                shadow=null
                moves=null
            }
        },
        function():void{
            var newRect = piece.getRect()
            var i=inGrid(w, h, newRect, game, gridPoints)
            if(i!=null){
                if(moves==null){
                    moves=game.canMove(w, h, cell)
                }
                if(-1 != moves.indexOf(i[0])){
                    var xy=game.toXY(i[0])
                    var newRectWidth=newRect.width()
                    var newRectHeight=newRect.height()
                    newRect.left=gridPoints[0][xy.x]+pieceMargin+ctx.lineWidth/2
                    newRect.top=gridPoints[1][xy.y]+pieceMargin+ctx.lineWidth/2
                    newRect.right=newRect.left+newRectWidth
                    newRect.bottom=newRect.top+newRectHeight
                    piece.setRect(newRect.left, newRect.top, newRect.right, newRect.bottom)
                    game.move(w, h, i[0], cell)
                    cell=i[0]
                    if(game.win(cell)){
                        createWinPanel(mng, bg, cc, game, canvas)
                    }
                }else{
                    newRect=pieceRect.copy()
                }
            }else{
                bg.remove(piece)
                game.remove(w, h, cell)
                if(sizeIdx==0){
                    tmpl.setDragable(true)
                    solveButton.Enable(false)
                }
            }
            pieceRect=newRect.copy()
            shadow=null
            piece.setRect(newRect.left, newRect.top, newRect.right, newRect.bottom)
            piece.setZ(pieceZ)
            moves=null
        },
        function():void{
            shadow=null
            piece.setRect(pieceRect.left, pieceRect.top, pieceRect.right, pieceRect.bottom)
            piece.setZ(pieceZ)
            moves=null
        }
    )
    game.put(w, h, cell, function(pos){
        cell=pos
    })
    bg.add(piece)
    game.piece(cell, piece)
}
function createWinPanel(mng:OperatableRectMng, bg:CanvasSprite, cc:CanvasContext, game:Board, canvas:any){
    var ctx:any=cc.ctx
    var uiLeft=0, uiTop=0, uiRight=canvas.width, uiBottom=canvas.height
    var text = "You win"
    var b=new Button(function(){
        bg.remove(bk)
        mng.remove(b)
        game.setDragable(true)
    }, function(){})
    b.Enable(true)
    b.setRect(uiLeft,uiTop, uiRight, uiBottom)
    b.setTransform(cc.getTransform())
    var bk=new CanvasButtonSprite(mng, cc.getTransform(), 10000, function(){
        ctx.fillStyle='rgb(128,128,128,0.95)'
        ctx.fillRect(uiLeft,uiTop, uiRight-uiLeft, uiBottom-uiTop)

        if(text){
            ctx.font = '36px Arial';
            ctx.fillStyle='rgb(0,0,0)'
            ctx.textAlign='center'
            ctx.textBaseline = 'middle'
            ctx.strokeStyle='rgb(255,255,255)'
            ctx.fillText(text, uiLeft+(uiRight-uiLeft)/2, 200)
        }
    }, b)
    bg.add(bk)
    b.setZ(bk.getZ())
    // mng.add(b)
}
function inGrid(w:number, h:number, r:CanvasRect, game:Board, gridPoints:Array<Array< number>>){
    var l=gridPoints[0].length
    var horIn:number|boolean=false, verIn:number|boolean=false
    for(var i=0; i<l-w; i++){
        if(gridPoints[0][i]<r.left && r.left<gridPoints[0][i+1]){
            horIn=i
            if(i!=l-w-1 && Math.abs(gridPoints[0][i+1]-r.left) < Math.abs(gridPoints[0][i]-r.left)){
                horIn=i+1
            }
            break
        }
    }
    if(horIn===false && (gridPoints[0][w-1]<r.right && r.right<gridPoints[0][w])){
        horIn=0
    }
    l=gridPoints[1].length
    for(var i=0; i<l-h; i++){
        if(gridPoints[1][i]<r.top && r.top<gridPoints[1][i+1]){
            verIn=i
            if(i!=l-h-1 && Math.abs(gridPoints[1][i+1]-r.top) < Math.abs(gridPoints[1][i]-r.top)){
                verIn=i+1
            }
            break
        }
    }
    if(verIn===false && (gridPoints[1][h-1]<r.bottom && r.bottom<gridPoints[1][h])){
        verIn=0
    }
    if(horIn!==false && verIn!==false){
        var ret = new Array()
        for(var i=0; i<w; i++){
            for(var j=0; j<h; j++){
                var v=game.toIndex(i+horIn, j+verIn)
                ret.push(v)
            }
        }
        return ret
    }else{
        return null
    }
}
function drawBorder(ctx, width, height){
    ctx.save()
    ctx.fillStyle='rgb(255,0,0,1)'
    ctx.strokeStyle='rgb(0, 0, 0, 1)'
    ctx.lineWidth=4;
    ctx.beginPath();
    ctx.arc (borderRadius+borderMargin+ctx.lineWidth/2, borderRadius+borderMargin+ctx.lineWidth/2, borderRadius, Math.PI*1.5, Math.PI, true)
    ctx.lineTo(borderMargin+ctx.lineWidth/2, height-borderRadius-borderMargin-ctx.lineWidth/2)
    ctx.arc(borderRadius+borderMargin+ctx.lineWidth/2, height-borderRadius-borderMargin-ctx.lineWidth/2, borderRadius, Math.PI, Math.PI/2, true)
    ctx.lineTo(width-2*borderMargin-ctx.lineWidth/2*2-2*borderRadius, height-borderMargin-ctx.lineWidth/2)
    ctx.arc(width-borderMargin-ctx.lineWidth/2-borderRadius, height-borderMargin-ctx.lineWidth/2-borderRadius, borderRadius, Math.PI/2, 0, true)
    ctx.lineTo(width-borderMargin-ctx.lineWidth/2, borderMargin+ctx.lineWidth/2+borderRadius)
    ctx.arc(width-borderMargin-ctx.lineWidth/2-borderRadius, borderMargin+ctx.lineWidth/2+borderRadius, borderRadius, 0, Math.PI*1.5, true)
    ctx.lineTo(borderMargin+ctx.lineWidth/2+borderRadius, borderMargin+ctx.lineWidth/2)
    ctx.stroke()
    ctx.fill()
    ctx.restore()
}
const dashLineWidth=2
function calculateGrids(cellWidth:number, cellHeight:number):Array<Array<number>>{
    var points = new Array(2)
    points[0] = new Array(horizon)
    points[1] = new Array(vertial)
    for(var i=0; i<=horizon; i++){
        points[0][i]=i*cellWidth+borderRadius+borderMargin+dashLineWidth/2
    }
    for(var i=0; i<=vertial; i++){
        points[1][i]=i*cellHeight+borderRadius+borderMargin+dashLineWidth/2;
    }
    return points
}
function drawGrid(ctx, width, height, gridPoints){
    ctx.save()
    ctx.strokeStyle = 'rgb(127, 127, 127, 0.5)'
    ctx.lineWidth=dashLineWidth
    ctx.setLineDash([4, 2])
    ctx.beginPath()
    for(var i=0; i<=horizon; i++){
        var x= gridPoints[0][i]
        ctx.moveTo(x, borderRadius+borderMargin+ctx.lineWidth/2)
        ctx.lineTo(x, height-borderMargin-borderRadius-ctx.lineWidth/2)
    }
    for(var i=0; i<=vertial; i++){
        var y=gridPoints[1][i]
        ctx.moveTo(borderRadius+borderMargin+ctx.lineWidth/2, y)
        ctx.lineTo(width-borderRadius-borderMargin-ctx.lineWidth/2, y)
    }
    ctx.stroke()
    ctx.restore()
}

function pieceDragRect(w, h, x, y, cellWidth, cellHeight){
    return new CanvasRect(x+pieceMargin, y+pieceMargin, x+w*cellWidth-2*pieceMargin-pieceWidth, y+h*cellHeight-2*pieceMargin-pieceWidth)
}
function drawPiece(w, h, ctx, x, y, lineColor, fillColor, cellWidth, cellHeight){
    ctx.lineWidth = pieceWidth
    ctx.strokeStyle=lineColor
    ctx.fillStyle=fillColor
    ctx.beginPath()
    ctx.arc(x+ctx.lineWidth/2+borderRadius, y+ctx.lineWidth/2+borderRadius, borderRadius, Math.PI*1.5, Math.PI, true)
    ctx.lineTo(x+ctx.lineWidth/2, y+h*cellHeight-2*pieceMargin-ctx.lineWidth/2-borderRadius)
    ctx.arc(x+ctx.lineWidth/2+borderRadius, y+h*cellHeight-2*pieceMargin-ctx.lineWidth/2-borderRadius, borderRadius, Math.PI, Math.PI/2, true)
    ctx.lineTo(x+w*cellWidth-2*pieceMargin-ctx.lineWidth/2-borderRadius, y+h*cellHeight-2*pieceMargin-ctx.lineWidth/2)
    ctx.arc(x+w*cellWidth-2*pieceMargin-ctx.lineWidth/2-borderRadius, y+h*cellHeight-2*pieceMargin-ctx.lineWidth/2-borderRadius, borderRadius, Math.PI/2, 0, true)
    ctx.lineTo(x+w*cellWidth-2*pieceMargin-ctx.lineWidth/2, y+ctx.lineWidth/2+borderRadius)
    ctx.arc(x+w*cellWidth-2*pieceMargin-ctx.lineWidth/2-borderRadius, y+ctx.lineWidth/2+borderRadius, borderRadius, 0, Math.PI*1.5, true)
    ctx.lineTo(x+ctx.lineWidth/2+borderRadius, y+ctx.lineWidth/2)
    ctx.fill()
    ctx.stroke()
}
