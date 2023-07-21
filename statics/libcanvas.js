var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (Object.prototype.hasOwnProperty.call(b, p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        if (typeof b !== "function" && b !== null)
            throw new TypeError("Class extends value " + String(b) + " is not a constructor or null");
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
var CanvasProperties = /** @class */ (function () {
    function CanvasProperties() {
    }
    CanvasProperties.prototype.checkTransform = function (ctx) {
        this.transformEnable = ctx.getTransform ? true : false;
    };
    CanvasProperties.prototype.getTransformEnable = function () {
        return this.transformEnable;
    };
    return CanvasProperties;
}());
var canvasProperties = new CanvasProperties();
var XYPair = /** @class */ (function () {
    function XYPair(x, y) {
        this.x = x;
        this.y = y;
    }
    return XYPair;
}());
function fromGlobalPoint(x, y, tr) {
    if (!canvasProperties.getTransformEnable()) {
        return new XYPair(x, y);
    }
    var a = new DOMMatrix(tr);
    a.invertSelf();
    return new XYPair(a.a * x + a.c * y + a.e, a.b * x + a.d * y + a.f);
}
function toGlobalPoint(x, y, tr) {
    return new XYPair(tr.a * x + tr.c * y + tr.e, tr.b * x + tr.d * y + tr.f);
}
var CanvasRect = /** @class */ (function () {
    function CanvasRect(left, top, right, bottom) {
        this.left = left;
        this.top = top;
        this.right = right;
        this.bottom = bottom;
    }
    CanvasRect.prototype.width = function () {
        return this.right - this.left;
    };
    CanvasRect.prototype.height = function () {
        return this.bottom - this.top;
    };
    CanvasRect.prototype.merge = function (a) {
        if (a == null) {
            return;
        }
        if (a.left < this.left) {
            this.left = a.left;
        }
        if (a.right > this.right) {
            this.right = a.right;
        }
        if (a.top < this.top) {
            this.top = a.top;
        }
        if (a.bottom > this.bottom) {
            this.bottom = a.bottom;
        }
    };
    CanvasRect.prototype.copy = function () {
        return new CanvasRect(this.left, this.top, this.right, this.bottom);
    };
    return CanvasRect;
}());
var minMouseMoveDist = 10;
var OperatableRectMng = /** @class */ (function () {
    function OperatableRectMng(target) {
        var _this = this;
        this.target = target;
        this.rects = new Array();
        this.moving = false;
        window.addEventListener("mousemove", function (ev) {
            if (_this.rectMouseDown != null) {
                if (ev.target == _this.target) {
                    if (_this.moving) {
                        _this.moveRect(ev);
                    }
                    else {
                        var dx = ev.offsetX - _this.pointMouseDown[0];
                        var dy = ev.offsetY - _this.pointMouseDown[1];
                        var dd = Math.sqrt(dx * dx + dy * dy);
                        if (dd > minMouseMoveDist && _this.rectMouseDown.isDragable()) {
                            _this.moving = true;
                            var pmd = fromGlobalPoint(_this.pointMouseDown[0], _this.pointMouseDown[1], _this.rectMouseDown.getTransform());
                            _this.pointMouseDown[0] = pmd.x;
                            _this.pointMouseDown[1] = pmd.y;
                            _this.moveRect(ev);
                        }
                    }
                }
                else if (_this.moving) {
                    _this.rectMouseDown.onOut();
                    _this.operationEnd();
                }
            }
        });
        window.addEventListener("mousedown", function (ev) {
            if (ev.target == _this.target) {
                var curRect;
                var curZ = Number.MIN_VALUE;
                var found;
                var pointMouseDown;
                _this.rects.forEach(function (rect) {
                    if (_this.inRect(rect, ev.offsetX, ev.offsetY)) {
                        var z = rect.getZ();
                        if (z >= curZ) {
                            curZ = z;
                            curRect = rect;
                            found = true;
                            pointMouseDown = [ev.offsetX, ev.offsetY];
                        }
                    }
                });
                if (found) {
                    _this.rectMouseDown = curRect;
                    _this.pointMouseDown = pointMouseDown;
                }
            }
        });
        window.addEventListener("mouseup", function (ev) {
            if (_this.rectMouseDown != null) {
                if (ev.target == _this.target) {
                    if (_this.moving) {
                        _this.rectMouseDown.onRelease();
                    }
                    else {
                        if (_this.inRect(_this.rectMouseDown, ev.offsetX, ev.offsetY)) {
                            _this.rectMouseDown.onClick();
                        }
                        else {
                            _this.rectMouseDown.onClickCancel();
                        }
                    }
                }
                else {
                    _this.rectMouseDown.onClickCancel();
                }
                _this.operationEnd();
            }
        });
    }
    OperatableRectMng.prototype.moveRect = function (ev) {
        var p = fromGlobalPoint(ev.offsetX, ev.offsetY, this.rectMouseDown.getTransform());
        this.rectMouseDown.onDrag(this.pointMouseDown[0], this.pointMouseDown[1], p.x, p.y);
    };
    OperatableRectMng.prototype.inRect = function (rect, x, y) {
        var p = fromGlobalPoint(x, y, rect.getTransform());
        var r = rect.getRect();
        return (r.top <= p.y && p.y <= r.bottom && r.left <= p.x && p.x <= r.right);
    };
    OperatableRectMng.prototype.add = function (rect) {
        this.rects.push(rect);
        rect.added();
    };
    OperatableRectMng.prototype.remove = function (rect) {
        var idx = this.rects.indexOf(rect);
        if (idx != -1) {
            this.rects.splice(idx, 1)[0].removed();
        }
    };
    OperatableRectMng.prototype.operationEnd = function () {
        this.rectMouseDown = null;
        this.pointMouseDown = null;
        this.moving = false;
    };
    return OperatableRectMng;
}());
var OperatableRect = /** @class */ (function () {
    function OperatableRect() {
        this.rect = new CanvasRect(0, 0, 0, 0);
        this.enable = false;
    }
    OperatableRect.prototype.added = function () { };
    OperatableRect.prototype.removed = function () { };
    OperatableRect.prototype.setTransform = function (tr) {
        this.tr = tr;
    };
    OperatableRect.prototype.getTransform = function () {
        return this.tr;
    };
    OperatableRect.prototype.setZ = function (z) {
        this.z = z;
    };
    OperatableRect.prototype.getZ = function () {
        return this.z;
    };
    OperatableRect.prototype.setRect = function (left, top, right, bottom) {
        this.rect.left = left;
        this.rect.top = top;
        this.rect.right = right;
        this.rect.bottom = bottom;
    };
    OperatableRect.prototype.getRect = function () {
        return this.rect;
    };
    OperatableRect.prototype.isDragable = function () {
        return this.dragable;
    };
    OperatableRect.prototype.setDragable = function (dragable) {
        this.dragable = dragable;
    };
    OperatableRect.prototype.Enable = function (enable) {
        this.enable = enable;
    };
    OperatableRect.prototype.onClick = function () { };
    OperatableRect.prototype.onClickCancel = function () { };
    OperatableRect.prototype.onDrag = function (downX, downY, x, y) { };
    OperatableRect.prototype.onRelease = function () { };
    OperatableRect.prototype.onOut = function () { };
    return OperatableRect;
}());
var Button = /** @class */ (function (_super) {
    __extends(Button, _super);
    function Button(clickFunc, cancelFunc) {
        var _this = _super.call(this) || this;
        _this.clickFunc = clickFunc;
        _this.cancelFunc = cancelFunc;
        return _this;
    }
    Button.prototype.onClick = function () {
        if (this.enable) {
            this.clickFunc();
        }
    };
    Button.prototype.onClickCancel = function () {
        this.cancelFunc();
    };
    return Button;
}(OperatableRect));
var DragableComponent = /** @class */ (function (_super) {
    __extends(DragableComponent, _super);
    function DragableComponent(onDrag, onRelease, onOut) {
        var _this = _super.call(this) || this;
        _this.onDragFunc = onDrag;
        _this.onReleaseFunc = onRelease;
        _this.onOutFunc = onOut;
        _this.dragable = true;
        return _this;
    }
    DragableComponent.prototype.onDrag = function (downX, downY, x, y) {
        this.onDragFunc(downX, downY, x, y);
    };
    DragableComponent.prototype.onRelease = function () {
        this.onReleaseFunc();
    };
    DragableComponent.prototype.onOut = function () {
        this.onOutFunc();
    };
    return DragableComponent;
}(OperatableRect));
var MovableComponent = /** @class */ (function (_super) {
    __extends(MovableComponent, _super);
    function MovableComponent() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    MovableComponent.prototype.setRect = function (left, top, right, bottom) {
        this.rawRect = new CanvasRect(0, 0, 0, 0);
        this.rawRect.left = left;
        this.rawRect.top = top;
        this.rawRect.right = right;
        this.rawRect.bottom = bottom;
        _super.prototype.setRect.call(this, left, top, right, bottom);
    };
    MovableComponent.prototype.onDrag = function (downX, downY, x, y) {
        var dx = downX - this.rawRect.left;
        var dy = downY - this.rawRect.top;
        _super.prototype.setRect.call(this, x - dx, y - dy, x - dx + this.rawRect.right - this.rawRect.left, y - dy + this.rawRect.bottom - this.rawRect.top);
        _super.prototype.onDrag.call(this, downX, downY, x, y);
    };
    MovableComponent.prototype.copyRect = function () {
        var r = _super.prototype.getRect.call(this);
        for (var index in r) {
            this.rawRect[index] = r[index];
        }
    };
    MovableComponent.prototype.setRecentRectWhenRelease = function (recentRectWhenRelease) {
        this.recentRectWhenRelease = recentRectWhenRelease;
    };
    MovableComponent.prototype.onRelease = function () {
        if (this.recentRectWhenRelease) {
            this.copyRect();
        }
        _super.prototype.onRelease.call(this);
    };
    MovableComponent.prototype.setRecentRectWhenOut = function (recentRectWhenOut) {
        this.recentRectWhenOut = recentRectWhenOut;
    };
    MovableComponent.prototype.onOut = function () {
        if (this.recentRectWhenOut) {
            this.copyRect();
        }
        _super.prototype.onOut.call(this);
    };
    return MovableComponent;
}(DragableComponent));
var CanvasSprite = /** @class */ (function () {
    function CanvasSprite(tr, z, paintFunc) {
        this.subs = new Array();
        this.tr = tr;
        this.z = z;
        this.paintFunc = paintFunc;
    }
    CanvasSprite.prototype.added = function () { };
    CanvasSprite.prototype.add = function (sub) {
        this.subs.push(sub);
        sub.added();
    };
    CanvasSprite.prototype.removed = function (level) {
        this.subs.forEach(function (s) {
            s.removed(level + 1);
        });
    };
    CanvasSprite.prototype.remove = function (sub) {
        sub.removed(0);
        var idx = this.subs.indexOf(sub);
        if (idx != -1) {
            this.subs.splice(idx, 1);
        }
    };
    CanvasSprite.prototype.setZ = function (z) {
        this.z = z;
    };
    CanvasSprite.prototype.getZ = function () { return this.z; };
    CanvasSprite.prototype.setTransform = function (tr) {
        this.tr = tr;
    };
    CanvasSprite.prototype.getTransform = function () {
        return this.tr;
    };
    CanvasSprite.prototype.refresh = function (ctx) {
        this.subs.sort(function (a, b) {
            var za = a.getZ();
            var zb = b.getZ();
            if (za < zb) {
                return -1;
            }
            else if (za > zb) {
                return +1;
            }
            return 0;
        });
        var r = new CanvasRect(0, 0, 0, 0);
        for (var idx in this.subs) {
            r.merge(this.subs[idx].range());
        }
        var selfOnPainted = false;
        for (var idx in this.subs) {
            if (!selfOnPainted && this.subs[idx].getZ() >= 0) {
                selfOnPainted = true;
                this.draw(ctx, r);
            }
            this.subs[idx].refresh(ctx);
        }
        if (!selfOnPainted) {
            this.draw(ctx, r);
        }
    };
    CanvasSprite.prototype.draw = function (ctx, r) {
        ctx.save();
        if (this.tr != null) {
            ctx.setTransform(this.tr);
        }
        this.paintFunc(ctx, r);
        ctx.restore();
    };
    CanvasSprite.prototype.range = function () {
        return null;
    };
    return CanvasSprite;
}());
function RootSprite(root, ctx) {
    var f = function () {
        root.refresh(ctx);
        window.requestAnimationFrame(f);
    };
    window.requestAnimationFrame(f);
}
var CanvasOperatableSprite = /** @class */ (function (_super) {
    __extends(CanvasOperatableSprite, _super);
    function CanvasOperatableSprite(mng, tr, z, paintFunc) {
        var _this = _super.call(this, tr, z, paintFunc) || this;
        _this.mng = mng;
        return _this;
    }
    CanvasOperatableSprite.prototype.setOperable = function (opRect) {
        this.opRect = opRect;
    };
    CanvasOperatableSprite.prototype.getOperable = function () {
        return this.opRect;
    };
    CanvasOperatableSprite.prototype.getRect = function () {
        return this.opRect.getRect();
    };
    CanvasOperatableSprite.prototype.setRect = function (left, top, right, bottom) {
        this.opRect.setRect(left, top, right, bottom);
    };
    CanvasOperatableSprite.prototype.setDragable = function (dragable) {
        this.opRect.setDragable(dragable);
    };
    CanvasOperatableSprite.prototype.removed = function (level) {
        _super.prototype.removed.call(this, level);
        this.mng.remove(this.getOperable());
    };
    CanvasOperatableSprite.prototype.add = function (sub) {
        _super.prototype.add.call(this, sub);
    };
    return CanvasOperatableSprite;
}(CanvasSprite));
var CanvasMovableSprite = /** @class */ (function (_super) {
    __extends(CanvasMovableSprite, _super);
    function CanvasMovableSprite(mng, tr, z, rect, paintFunc, onDrag, onRelease, onOut) {
        var _this = _super.call(this, mng, tr, z, paintFunc) || this;
        var mc = new MovableComponent(onDrag, onRelease, onOut);
        mc.setTransform(tr);
        mc.setZ(z);
        mc.setRect(rect.left, rect.top, rect.right, rect.bottom);
        mc.setRecentRectWhenOut(true);
        _this.setOperable(mc);
        mng.add(mc);
        _this.mng = mng;
        return _this;
    }
    CanvasMovableSprite.prototype.getOperable = function () {
        return this.opRect;
    };
    return CanvasMovableSprite;
}(CanvasOperatableSprite));
var CanvasButtonSprite = /** @class */ (function (_super) {
    __extends(CanvasButtonSprite, _super);
    function CanvasButtonSprite(mng, tr, z, paintFunc, b) {
        var _this = _super.call(this, mng, tr, z, paintFunc) || this;
        _this.setOperable(b);
        mng.add(b);
        return _this;
    }
    return CanvasButtonSprite;
}(CanvasOperatableSprite));
