class CanvasProperties{
    transformEnable:boolean
    checkTransform(ctx:any){
        this.transformEnable=ctx.getTransform ? true : false
    }
    getTransformEnable():boolean{
        return this.transformEnable
    }
}
var canvasProperties=new CanvasProperties()

class XYPair{
    constructor(x:number, y:number){
        this.x=x
        this.y=y
    }
    x:number
    y:number
}
function fromGlobalPoint(x:number, y:number, tr:any):XYPair{
    if( !canvasProperties.getTransformEnable() ){
        return new XYPair(x, y)
    }
    var a = new DOMMatrix(tr)
    a.invertSelf()
    return new XYPair(a.a * x + a.c * y + a.e, a.b * x + a.d * y + a.f)
}
function toGlobalPoint(x:number, y:number, tr:any):XYPair{
    return new XYPair(tr.a * x + tr.c * y + tr.e, tr.b * x + tr.d * y + tr.f)
}
class CanvasRect{
    left:number
    top:number
    right:number
    bottom:number
    constructor(left:number,top:number,right:number,bottom:number){
        this.left=left
        this.top=top
        this.right=right
        this.bottom=bottom
    }
    width():number{
        return this.right-this.left
    }
    height():number{
        return this.bottom-this.top
    }
    merge(a:CanvasRect):void{
        if(a==null){
            return
        }
        if(a.left<this.left){
            this.left=a.left
        }
        if(a.right>this.right){
            this.right=a.right
        }
        if(a.top<this.top){
            this.top=a.top
        }
        if(a.bottom>this.bottom){
            this.bottom=a.bottom
        }
    }
    copy():CanvasRect{
        return new CanvasRect(this.left, this.top, this.right, this.bottom)
    }
}
const minMouseMoveDist=10
class OperatableRectMng{
    rects:Array<OperatableRect>
    rectMouseDown:OperatableRect
    pointMouseDown:Array<number>
    target:any
    moving:boolean
    constructor(target:any){
        this.target=target
        this.rects = new Array()
        this.moving=false
        window.addEventListener("mousemove",(ev:any):void=>{
            if(this.rectMouseDown != null){
                if(ev.target==this.target){
                    if(this.moving){
                        this.moveRect(ev)
                    }else{
                        var dx=ev.offsetX-this.pointMouseDown[0]
                        var dy=ev.offsetY-this.pointMouseDown[1]
                        var dd=Math.sqrt(dx*dx+dy*dy)
                        if(dd>minMouseMoveDist && this.rectMouseDown.isDragable()){
                            this.moving=true
                            var pmd = fromGlobalPoint(this.pointMouseDown[0], this.pointMouseDown[1], this.rectMouseDown.getTransform())
                            this.pointMouseDown[0]=pmd.x
                            this.pointMouseDown[1]=pmd.y
                            this.moveRect(ev)
                        }
                    }
                }else if(this.moving){
                    this.rectMouseDown.onOut()
                    this.operationEnd()
                }
            }
        })
        window.addEventListener("mousedown", (ev:any):void=>{
            if(ev.target==this.target){
                var curRect:OperatableRect
                var curZ=Number.MIN_VALUE
                var found:boolean
                var pointMouseDown:number[]
                this.rects.forEach(rect=>{
                    if (this.inRect(rect, ev.offsetX, ev.offsetY)){
                        var z=rect.getZ()
                        if(z>=curZ){
                            curZ=z
                            curRect=rect
                            found=true
                            pointMouseDown=[ev.offsetX, ev.offsetY]
                        }
                    }
                })
                if(found){
                    this.rectMouseDown=curRect
                    this.pointMouseDown=pointMouseDown
                }
            }
        })
        window.addEventListener("mouseup", (ev:any):void=>{
            if(this.rectMouseDown!=null){
                if(ev.target==this.target){
                    if(this.moving){
                        this.rectMouseDown.onRelease()
                    }else{
                        if(this.inRect(this.rectMouseDown, ev.offsetX, ev.offsetY)){
                            this.rectMouseDown.onClick()
                        }else{
                            this.rectMouseDown.onClickCancel()
                        }
                    }
                }else{
                    this.rectMouseDown.onClickCancel()
                }
                this.operationEnd()
            }
        })
    }
    moveRect(ev:any):void{
        var p=fromGlobalPoint(ev.offsetX, ev.offsetY, this.rectMouseDown.getTransform())
        this.rectMouseDown.onDrag(this.pointMouseDown[0], this.pointMouseDown[1], p.x,p.y)
    }
    inRect(rect:OperatableRect, x:number, y:number):boolean{
        var p=fromGlobalPoint(x, y, rect.getTransform())
        var r=rect.getRect()
        return (r.top<=p.y && p.y<=r.bottom && r.left<=p.x && p.x<=r.right)
    }
    add(rect:OperatableRect):void{
        this.rects.push(rect)
        rect.added()
    }
    remove(rect:OperatableRect):void{
        var idx=this.rects.indexOf(rect)
        if(idx!=-1){
            this.rects.splice(idx,1)[0].removed()
        }
    }
    operationEnd():void{
        this.rectMouseDown=null
        this.pointMouseDown=null
        this.moving=false
    }
}
class OperatableRect {
    rect:CanvasRect
    constructor(){
        this.rect=new CanvasRect(0,0,0,0)
        this.enable=false
    }
    added():void{}
    removed():void{}
    tr:any
    setTransform(tr:any):void{
        this.tr = tr
    }
    getTransform():any{
        return this.tr
    }
    z:number
    setZ(z:number):void{
        this.z=z
    }
    getZ():number{
        return this.z
    }
    setRect(left:number, top:number, right:number, bottom:number):void{
        this.rect.left=left
        this.rect.top=top
        this.rect.right=right
        this.rect.bottom=bottom
    }
    getRect():CanvasRect{
        return this.rect
    }
    dragable:boolean
    isDragable():boolean{
        return this.dragable
    }
    setDragable(dragable:boolean):void{
        this.dragable=dragable
    }
    enable:boolean
    Enable(enable:boolean):void{
        this.enable=enable
    }
    onClick():void{}
    onClickCancel():void{}
    onDrag(downX:number, downY:number, x:number, y:number):void{}
    onRelease():void{}
    onOut():void{}
}

class Button extends OperatableRect{
    clickFunc:()=>void
    cancelFunc:()=>void
    constructor(clickFunc:()=>void, cancelFunc:()=>void){
        super()
        this.clickFunc=clickFunc
        this.cancelFunc=cancelFunc
    }
    onClick():void{
        if(this.enable){
            this.clickFunc()
        }
    }
    onClickCancel():void{
        this.cancelFunc()
    }
}

class DragableComponent extends OperatableRect{
    onDragFunc:(downX:number, downY:number, x:number, y:number)=>void
    onReleaseFunc:()=>void
    onOutFunc:()=>void
    constructor(onDrag:(downX:number, downY:number, x:number, y:number)=>void, onRelease:()=>void, onOut:()=>void){
        super()
        this.onDragFunc=onDrag
        this.onReleaseFunc=onRelease
        this.onOutFunc=onOut
        this.dragable=true
    }
    onDrag(downX:number, downY:number, x:number, y:number):void{
        this.onDragFunc(downX, downY, x, y)
    }
    onRelease():void{
        this.onReleaseFunc()
    }
    onOut():void{
        this.onOutFunc()
    }
}
class MovableComponent extends DragableComponent{
    rawRect:CanvasRect
    setRect(left:number, top:number, right:number, bottom:number):void{
        this.rawRect=new CanvasRect(0,0,0,0)
        this.rawRect.left=left
        this.rawRect.top=top
        this.rawRect.right=right
        this.rawRect.bottom=bottom
        super.setRect(left, top, right, bottom)
    }
    onDrag(downX:number, downY:number, x:number, y:number):void{
        var dx=downX-this.rawRect.left
        var dy=downY-this.rawRect.top
        super.setRect(x-dx, y-dy, x-dx+this.rawRect.right-this.rawRect.left, y-dy+this.rawRect.bottom-this.rawRect.top)
        super.onDrag(downX, downY, x, y)
    }
    copyRect():void{
        var r=super.getRect()
        for (let index in r){
            this.rawRect[index]=r[index]
        }
    }
    recentRectWhenRelease:boolean
    setRecentRectWhenRelease(recentRectWhenRelease:boolean):void{
        this.recentRectWhenRelease=recentRectWhenRelease
    }
    onRelease():void{
        if(this.recentRectWhenRelease){
            this.copyRect()
        }
        super.onRelease()
    }
    recentRectWhenOut:boolean
    setRecentRectWhenOut(recentRectWhenOut:boolean):void{
        this.recentRectWhenOut=recentRectWhenOut
    }
    onOut():void{
        if(this.recentRectWhenOut){
            this.copyRect()
        }
        super.onOut()
    }
}

class CanvasSprite{
    subs:Array<CanvasSprite>
    paintFunc:(any,CanvasRect)=>void
    constructor(tr:any, z:number, paintFunc:(any,CanvasRect)=>void){
        this.subs = new Array()
        this.tr=tr
        this.z =z
        this.paintFunc=paintFunc
    }
    added():void{}
    add(sub:CanvasSprite):void{
        this.subs.push(sub)
        sub.added()
    }
    removed(level:number):void{
        this.subs.forEach((s:CanvasSprite):void=>{
            s.removed(level+1)
        })
    }
    remove(sub:CanvasSprite):void{
        sub.removed(0)
        var idx=this.subs.indexOf(sub)
        if(idx!=-1){
            this.subs.splice(idx, 1)
        }
    }
    z:number
    setZ(z:number):void{
        this.z=z
    }
    getZ():number{return this.z}
    tr:any
    setTransform(tr:any):void{
        this.tr=tr
    }
    getTransform():any{
        return this.tr
    }
    refresh(ctx:any):void{
        this.subs.sort(function(a:CanvasSprite, b:CanvasSprite){
            let za=a.getZ()
            let zb=b.getZ()
            if(za<zb){
                return -1
            }else if (za>zb){
                return +1
            }
            return 0
        })
        var r = new CanvasRect(0,0,0,0)
        for(let idx in this.subs){
            r.merge(this.subs[idx].range())
        }
        var selfOnPainted = false
        for(let idx in this.subs){
            if(!selfOnPainted && this.subs[idx].getZ()>=0){
                selfOnPainted=true
                this.draw(ctx, r)
            }
            this.subs[idx].refresh(ctx)
        }
        if(!selfOnPainted){
            this.draw(ctx, r)
        }
    }
    draw(ctx:any, r:CanvasRect):void{
        ctx.save()
        if(this.tr != null){
                ctx.setTransform(this.tr)
        }
        this.paintFunc(ctx, r)
        ctx.restore()
    }
    range():CanvasRect{
        return null
    }
}
function RootSprite(root:any, ctx:any):void{
    let f= function():void{
        root.refresh(ctx)
        window.requestAnimationFrame(f)
    }
    window.requestAnimationFrame(f)
}

class CanvasOperatableSprite extends CanvasSprite{
    mng:OperatableRectMng
    opRect:OperatableRect
    constructor(mng:OperatableRectMng, tr:any, z:number, paintFunc:(any,CanvasRect)=>void){
        super( tr, z, paintFunc)
        this.mng=mng
    }
    setOperable(opRect:OperatableRect):void{
        this.opRect=opRect
    }
    getOperable():OperatableRect{
        return this.opRect
    }
    getRect():CanvasRect{
        return this.opRect.getRect()
    }
    setRect(left:number, top:number, right:number, bottom:number):void{
        this.opRect.setRect(left, top, right, bottom)
    }
    setDragable(dragable:boolean):void{
        this.opRect.setDragable(dragable)
    }
    override removed(level:number):void{
        super.removed(level)
        this.mng.remove(this.getOperable())
    }
    override add(sub:CanvasSprite):void{
        super.add(sub)
    }
}
class CanvasMovableSprite extends CanvasOperatableSprite{
    constructor(mng:OperatableRectMng, tr:any, z:number, rect:CanvasRect, paintFunc:(any,CanvasRect)=>void, onDrag:(downX:number, downY:number, x:number, y:number)=>void, onRelease:()=>void, onOut:()=>void){
        super(mng, tr, z, paintFunc)
        let mc=new MovableComponent(onDrag, onRelease, onOut)
        mc.setTransform(tr)
        mc.setZ(z)
        mc.setRect(rect.left, rect.top, rect.right, rect.bottom)
        mc.setRecentRectWhenOut(true)
        this.setOperable(mc)
        mng.add(mc)
        this.mng=mng
    }
    getOperable():MovableComponent{
        return this.opRect as MovableComponent
    }
}
class CanvasButtonSprite extends CanvasOperatableSprite{
    constructor(mng:OperatableRectMng, tr:any, z:number, paintFunc:(any,CanvasRect)=>void, b:Button){
        super(mng, tr, z, paintFunc)
        this.setOperable(b)
        mng.add(b)
    }
}