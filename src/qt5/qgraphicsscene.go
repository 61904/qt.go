package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtWidgets/qgraphicsscene.h
// dst-file: /src/widgets/qgraphicsscene.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QGraphicsScene::setForegroundBrush(const QBrush & brush);
extern void _ZN14QGraphicsScene18setForegroundBrushERK6QBrush(void* qthis, void* arg0);
  // proto:  void QGraphicsScene::setSceneRect(const QRectF & rect);
extern void _ZN14QGraphicsScene12setSceneRectERK6QRectF(void* qthis, void* arg0);
  // proto:  bool QGraphicsScene::isActive();
extern void _ZNK14QGraphicsScene8isActiveEv(void* qthis);
  // proto:  bool QGraphicsScene::hasFocus();
extern void _ZNK14QGraphicsScene8hasFocusEv(void* qthis);
  // proto:  QRectF QGraphicsScene::itemsBoundingRect();
extern void _ZNK14QGraphicsScene17itemsBoundingRectEv(void* qthis);
  // proto:  bool QGraphicsScene::sendEvent(QGraphicsItem * item, QEvent * event);
extern void _ZN14QGraphicsScene9sendEventEP13QGraphicsItemP6QEvent(void* qthis, void* arg0, void* arg1);
  // proto:  qreal QGraphicsScene::minimumRenderSize();
extern void _ZNK14QGraphicsScene17minimumRenderSizeEv(void* qthis);
  // proto:  QPainterPath QGraphicsScene::selectionArea();
extern void _ZNK14QGraphicsScene13selectionAreaEv(void* qthis);
  // proto:  void QGraphicsScene::update(const QRectF & rect);
extern void _ZN14QGraphicsScene6updateERK6QRectF(void* qthis, void* arg0);
  // proto:  QGraphicsPolygonItem * QGraphicsScene::addPolygon(const QPolygonF & polygon, const QPen & pen, const QBrush & brush);
extern void _ZN14QGraphicsScene10addPolygonERK9QPolygonFRK4QPenRK6QBrush(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  QGraphicsLineItem * QGraphicsScene::addLine(const QLineF & line, const QPen & pen);
extern void _ZN14QGraphicsScene7addLineERK6QLineFRK4QPen(void* qthis, void* arg0, void* arg1);
  // proto:  QPalette QGraphicsScene::palette();
extern void _ZNK14QGraphicsScene7paletteEv(void* qthis);
  // proto:  bool QGraphicsScene::isSortCacheEnabled();
extern void _ZNK14QGraphicsScene18isSortCacheEnabledEv(void* qthis);
  // proto:  void QGraphicsScene::QGraphicsScene(const QRectF & sceneRect, QObject * parent);
extern void* dector_ZN14QGraphicsSceneC1ERK6QRectFP7QObject(void* arg0, void* arg1);
extern void _ZN14QGraphicsSceneC1ERK6QRectFP7QObject(void* qthis, void* arg0, void* arg1);
  // proto:  void QGraphicsScene::QGraphicsScene(QObject * parent);
extern void* dector_ZN14QGraphicsSceneC1EP7QObject(void* arg0);
extern void _ZN14QGraphicsSceneC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QGraphicsScene::clearFocus();
extern void _ZN14QGraphicsScene10clearFocusEv(void* qthis);
  // proto:  const QMetaObject * QGraphicsScene::metaObject();
extern void _ZNK14QGraphicsScene10metaObjectEv(void* qthis);
  // proto:  QGraphicsSimpleTextItem * QGraphicsScene::addSimpleText(const QString & text, const QFont & font);
extern void _ZN14QGraphicsScene13addSimpleTextERK7QStringRK5QFont(void* qthis, void* arg0, void* arg1);
  // proto:  QGraphicsLineItem * QGraphicsScene::addLine(qreal x1, qreal y1, qreal x2, qreal y2, const QPen & pen);
extern void demth_ZN14QGraphicsScene7addLineEddddRK4QPen(void* qthis, double arg0, double arg1, double arg2, double arg3, void* arg4);
  // proto:  void QGraphicsScene::setBspTreeDepth(int depth);
extern void _ZN14QGraphicsScene15setBspTreeDepthEi(void* qthis, int32_t arg0);
  // proto:  QRectF QGraphicsScene::sceneRect();
extern void _ZNK14QGraphicsScene9sceneRectEv(void* qthis);
  // proto:  QGraphicsWidget * QGraphicsScene::activeWindow();
extern void _ZNK14QGraphicsScene12activeWindowEv(void* qthis);
  // proto:  QBrush QGraphicsScene::backgroundBrush();
extern void _ZNK14QGraphicsScene15backgroundBrushEv(void* qthis);
  // proto:  QGraphicsItem * QGraphicsScene::itemAt(qreal x, qreal y, const QTransform & deviceTransform);
extern void demth_ZNK14QGraphicsScene6itemAtEddRK10QTransform(void* qthis, double arg0, double arg1, void* arg2);
  // proto:  void QGraphicsScene::advance();
extern void _ZN14QGraphicsScene7advanceEv(void* qthis);
  // proto:  void QGraphicsScene::setStickyFocus(bool enabled);
extern void _ZN14QGraphicsScene14setStickyFocusEb(void* qthis, bool arg0);
  // proto:  QList<QGraphicsItem *> QGraphicsScene::selectedItems();
extern void _ZNK14QGraphicsScene13selectedItemsEv(void* qthis);
  // proto:  void QGraphicsScene::clear();
extern void _ZN14QGraphicsScene5clearEv(void* qthis);
  // proto:  void QGraphicsScene::setActivePanel(QGraphicsItem * item);
extern void _ZN14QGraphicsScene14setActivePanelEP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  QGraphicsPixmapItem * QGraphicsScene::addPixmap(const QPixmap & pixmap);
extern void _ZN14QGraphicsScene9addPixmapERK7QPixmap(void* qthis, void* arg0);
  // proto:  QBrush QGraphicsScene::foregroundBrush();
extern void _ZNK14QGraphicsScene15foregroundBrushEv(void* qthis);
  // proto:  QList<QGraphicsView *> QGraphicsScene::views();
extern void _ZNK14QGraphicsScene5viewsEv(void* qthis);
  // proto:  void QGraphicsScene::~QGraphicsScene();
extern void _ZN14QGraphicsSceneD0Ev(void* qthis);
  // proto:  QGraphicsRectItem * QGraphicsScene::addRect(qreal x, qreal y, qreal w, qreal h, const QPen & pen, const QBrush & brush);
extern void demth_ZN14QGraphicsScene7addRectEddddRK4QPenRK6QBrush(void* qthis, double arg0, double arg1, double arg2, double arg3, void* arg4, void* arg5);
  // proto:  int QGraphicsScene::bspTreeDepth();
extern void _ZNK14QGraphicsScene12bspTreeDepthEv(void* qthis);
  // proto:  void QGraphicsScene::setSceneRect(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZN14QGraphicsScene12setSceneRectEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QGraphicsScene::setStyle(QStyle * style);
extern void _ZN14QGraphicsScene8setStyleEP6QStyle(void* qthis, void* arg0);
  // proto:  void QGraphicsScene::setPalette(const QPalette & palette);
extern void _ZN14QGraphicsScene10setPaletteERK8QPalette(void* qthis, void* arg0);
  // proto:  void QGraphicsScene::setMinimumRenderSize(qreal minSize);
extern void _ZN14QGraphicsScene20setMinimumRenderSizeEd(void* qthis, double arg0);
  // proto:  void QGraphicsScene::QGraphicsScene(qreal x, qreal y, qreal width, qreal height, QObject * parent);
extern void* dector_ZN14QGraphicsSceneC1EddddP7QObject(double arg0, double arg1, double arg2, double arg3, void* arg4);
extern void _ZN14QGraphicsSceneC1EddddP7QObject(void* qthis, double arg0, double arg1, double arg2, double arg3, void* arg4);
  // proto:  QGraphicsItem * QGraphicsScene::mouseGrabberItem();
extern void _ZNK14QGraphicsScene16mouseGrabberItemEv(void* qthis);
  // proto:  QGraphicsRectItem * QGraphicsScene::addRect(const QRectF & rect, const QPen & pen, const QBrush & brush);
extern void _ZN14QGraphicsScene7addRectERK6QRectFRK4QPenRK6QBrush(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  QGraphicsEllipseItem * QGraphicsScene::addEllipse(const QRectF & rect, const QPen & pen, const QBrush & brush);
extern void _ZN14QGraphicsScene10addEllipseERK6QRectFRK4QPenRK6QBrush(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  qreal QGraphicsScene::height();
extern void demth_ZNK14QGraphicsScene6heightEv(void* qthis);
  // proto:  void QGraphicsScene::setSelectionArea(const QPainterPath & path, const QTransform & deviceTransform);
extern void _ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathRK10QTransform(void* qthis, void* arg0, void* arg1);
  // proto:  QFont QGraphicsScene::font();
extern void _ZNK14QGraphicsScene4fontEv(void* qthis);
  // proto:  void QGraphicsScene::clearSelection();
extern void _ZN14QGraphicsScene14clearSelectionEv(void* qthis);
  // proto:  void QGraphicsScene::QGraphicsScene(const QGraphicsScene & );
extern void* dector_ZN14QGraphicsSceneC1ERKS_(void* arg0);
extern void _ZN14QGraphicsSceneC1ERKS_(void* qthis, void* arg0);
  // proto:  void QGraphicsScene::removeItem(QGraphicsItem * item);
extern void _ZN14QGraphicsScene10removeItemEP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  QGraphicsEllipseItem * QGraphicsScene::addEllipse(qreal x, qreal y, qreal w, qreal h, const QPen & pen, const QBrush & brush);
extern void demth_ZN14QGraphicsScene10addEllipseEddddRK4QPenRK6QBrush(void* qthis, double arg0, double arg1, double arg2, double arg3, void* arg4, void* arg5);
  // proto:  void QGraphicsScene::setActiveWindow(QGraphicsWidget * widget);
extern void _ZN14QGraphicsScene15setActiveWindowEP15QGraphicsWidget(void* qthis, void* arg0);
  // proto:  QGraphicsItem * QGraphicsScene::focusItem();
extern void _ZNK14QGraphicsScene9focusItemEv(void* qthis);
  // proto:  QGraphicsTextItem * QGraphicsScene::addText(const QString & text, const QFont & font);
extern void _ZN14QGraphicsScene7addTextERK7QStringRK5QFont(void* qthis, void* arg0, void* arg1);
  // proto:  void QGraphicsScene::setSortCacheEnabled(bool enabled);
extern void _ZN14QGraphicsScene19setSortCacheEnabledEb(void* qthis, bool arg0);
  // proto:  QGraphicsItem * QGraphicsScene::itemAt(const QPointF & pos, const QTransform & deviceTransform);
extern void _ZNK14QGraphicsScene6itemAtERK7QPointFRK10QTransform(void* qthis, void* arg0, void* arg1);
  // proto:  void QGraphicsScene::destroyItemGroup(QGraphicsItemGroup * group);
extern void _ZN14QGraphicsScene16destroyItemGroupEP18QGraphicsItemGroup(void* qthis, void* arg0);
  // proto:  qreal QGraphicsScene::width();
extern void demth_ZNK14QGraphicsScene5widthEv(void* qthis);
  // proto:  void QGraphicsScene::update(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZN14QGraphicsScene6updateEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QGraphicsScene::addItem(QGraphicsItem * item);
extern void _ZN14QGraphicsScene7addItemEP13QGraphicsItem(void* qthis, void* arg0);
  // proto:  void QGraphicsScene::setBackgroundBrush(const QBrush & brush);
extern void _ZN14QGraphicsScene18setBackgroundBrushERK6QBrush(void* qthis, void* arg0);
  // proto:  QGraphicsItem * QGraphicsScene::activePanel();
extern void _ZNK14QGraphicsScene11activePanelEv(void* qthis);
  // proto:  QStyle * QGraphicsScene::style();
extern void _ZNK14QGraphicsScene5styleEv(void* qthis);
  // proto:  void QGraphicsScene::setFont(const QFont & font);
extern void _ZN14QGraphicsScene7setFontERK5QFont(void* qthis, void* arg0);
  // proto:  QGraphicsPathItem * QGraphicsScene::addPath(const QPainterPath & path, const QPen & pen, const QBrush & brush);
extern void _ZN14QGraphicsScene7addPathERK12QPainterPathRK4QPenRK6QBrush(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  bool QGraphicsScene::stickyFocus();
extern void _ZNK14QGraphicsScene11stickyFocusEv(void* qthis);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QGraphicsScene)=1
type QGraphicsScene struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _changed QGraphicsScene_changed_signal;
//  _sceneRectChanged QGraphicsScene_sceneRectChanged_signal;
//  _selectionChanged QGraphicsScene_selectionChanged_signal;
//  _focusItemChanged QGraphicsScene_focusItemChanged_signal;
}

  // proto:  void QGraphicsScene::setForegroundBrush(const QBrush & brush);
func (this *QGraphicsScene) setForegroundBrush(args ...interface{}) () {
  // setForegroundBrush(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene18setForegroundBrushERK6QBrush
    // invoke: void setForegroundBrush(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScene18setForegroundBrushERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setForegroundBrush", args)
  }

}

  // proto:  void QGraphicsScene::setSceneRect(const QRectF & rect);
func (this *QGraphicsScene) setSceneRect(args ...interface{}) () {
  // setSceneRect(const class QRectF &)
  // setSceneRect(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene12setSceneRectERK6QRectF
    // invoke: void setSceneRect(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScene12setSceneRectERK6QRectF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN14QGraphicsScene12setSceneRectEdddd
    // invoke: void setSceneRect(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.demth_ZN14QGraphicsScene12setSceneRectEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setSceneRect", args)
  }

}

  // proto:  bool QGraphicsScene::isActive();
func (this *QGraphicsScene) isActive(args ...interface{}) () {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene8isActiveEv
    // invoke: bool isActive()
    C._ZNK14QGraphicsScene8isActiveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "isActive", args)
  }

}

  // proto:  bool QGraphicsScene::hasFocus();
func (this *QGraphicsScene) hasFocus(args ...interface{}) () {
  // hasFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene8hasFocusEv
    // invoke: bool hasFocus()
    C._ZNK14QGraphicsScene8hasFocusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "hasFocus", args)
  }

}

  // proto:  QRectF QGraphicsScene::itemsBoundingRect();
func (this *QGraphicsScene) itemsBoundingRect(args ...interface{}) () {
  // itemsBoundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene17itemsBoundingRectEv
    // invoke: QRectF itemsBoundingRect()
    C._ZNK14QGraphicsScene17itemsBoundingRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "itemsBoundingRect", args)
  }

}

  // proto:  bool QGraphicsScene::sendEvent(QGraphicsItem * item, QEvent * event);
func (this *QGraphicsScene) sendEvent(args ...interface{}) () {
  // sendEvent(class QGraphicsItem *, class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"
  vtys[0][1] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene9sendEventEP13QGraphicsItemP6QEvent
    // invoke: bool sendEvent(class QGraphicsItem *, class QEvent *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QEvent).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN14QGraphicsScene9sendEventEP13QGraphicsItemP6QEvent(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "sendEvent", args)
  }

}

  // proto:  qreal QGraphicsScene::minimumRenderSize();
func (this *QGraphicsScene) minimumRenderSize(args ...interface{}) () {
  // minimumRenderSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene17minimumRenderSizeEv
    // invoke: qreal minimumRenderSize()
    C._ZNK14QGraphicsScene17minimumRenderSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "minimumRenderSize", args)
  }

}

  // proto:  QPainterPath QGraphicsScene::selectionArea();
func (this *QGraphicsScene) selectionArea(args ...interface{}) () {
  // selectionArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene13selectionAreaEv
    // invoke: QPainterPath selectionArea()
    C._ZNK14QGraphicsScene13selectionAreaEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "selectionArea", args)
  }

}

  // proto:  void QGraphicsScene::update(const QRectF & rect);
func (this *QGraphicsScene) update(args ...interface{}) () {
  // update(const class QRectF &)
  // update(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene6updateERK6QRectF
    // invoke: void update(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScene6updateERK6QRectF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN14QGraphicsScene6updateEdddd
    // invoke: void update(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.demth_ZN14QGraphicsScene6updateEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "update", args)
  }

}

  // proto:  QGraphicsPolygonItem * QGraphicsScene::addPolygon(const QPolygonF & polygon, const QPen & pen, const QBrush & brush);
func (this *QGraphicsScene) addPolygon(args ...interface{}) () {
  // addPolygon(const class QPolygonF &, const class QPen &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[0][1] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[0][2] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10addPolygonERK9QPolygonFRK4QPenRK6QBrush
    // invoke: QGraphicsPolygonItem * addPolygon(const class QPolygonF &, const class QPen &, const class QBrush &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPen).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QBrush).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN14QGraphicsScene10addPolygonERK9QPolygonFRK4QPenRK6QBrush(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addPolygon", args)
  }

}

  // proto:  QGraphicsLineItem * QGraphicsScene::addLine(const QLineF & line, const QPen & pen);
func (this *QGraphicsScene) addLine(args ...interface{}) () {
  // addLine(const class QLineF &, const class QPen &)
  // addLine(qreal, qreal, qreal, qreal, const class QPen &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"
  vtys[0][1] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = reflect.TypeOf(QPen{}) // "const QPen &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addLineERK6QLineFRK4QPen
    // invoke: QGraphicsLineItem * addLine(const class QLineF &, const class QPen &)
    var arg0 = args[0].(QLineF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPen).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN14QGraphicsScene7addLineERK6QLineFRK4QPen(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN14QGraphicsScene7addLineEddddRK4QPen
    // invoke: QGraphicsLineItem * addLine(qreal, qreal, qreal, qreal, const class QPen &)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QPen).qclsinst
    if false {fmt.Println(arg4)}
    C.demth_ZN14QGraphicsScene7addLineEddddRK4QPen(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addLine", args)
  }

}

  // proto:  QPalette QGraphicsScene::palette();
func (this *QGraphicsScene) palette(args ...interface{}) () {
  // palette()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene7paletteEv
    // invoke: QPalette palette()
    C._ZNK14QGraphicsScene7paletteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "palette", args)
  }

}

  // proto:  bool QGraphicsScene::isSortCacheEnabled();
func (this *QGraphicsScene) isSortCacheEnabled(args ...interface{}) () {
  // isSortCacheEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene18isSortCacheEnabledEv
    // invoke: bool isSortCacheEnabled()
    C._ZNK14QGraphicsScene18isSortCacheEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "isSortCacheEnabled", args)
  }

}

  // proto:  void QGraphicsScene::QGraphicsScene(const QRectF & sceneRect, QObject * parent);
func NewQGraphicsScene(args ...interface{}) QGraphicsScene {
  return QGraphicsScene{}
}

  // proto:  void QGraphicsScene::clearFocus();
func (this *QGraphicsScene) clearFocus(args ...interface{}) () {
  // clearFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10clearFocusEv
    // invoke: void clearFocus()
    C._ZN14QGraphicsScene10clearFocusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "clearFocus", args)
  }

}

  // proto:  const QMetaObject * QGraphicsScene::metaObject();
func (this *QGraphicsScene) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK14QGraphicsScene10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "metaObject", args)
  }

}

  // proto:  QGraphicsSimpleTextItem * QGraphicsScene::addSimpleText(const QString & text, const QFont & font);
func (this *QGraphicsScene) addSimpleText(args ...interface{}) () {
  // addSimpleText(const class QString &, const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene13addSimpleTextERK7QStringRK5QFont
    // invoke: QGraphicsSimpleTextItem * addSimpleText(const class QString &, const class QFont &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QFont).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN14QGraphicsScene13addSimpleTextERK7QStringRK5QFont(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addSimpleText", args)
  }

}

  // proto:  void QGraphicsScene::setBspTreeDepth(int depth);
func (this *QGraphicsScene) setBspTreeDepth(args ...interface{}) () {
  // setBspTreeDepth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene15setBspTreeDepthEi
    // invoke: void setBspTreeDepth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScene15setBspTreeDepthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setBspTreeDepth", args)
  }

}

  // proto:  QRectF QGraphicsScene::sceneRect();
func (this *QGraphicsScene) sceneRect(args ...interface{}) () {
  // sceneRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene9sceneRectEv
    // invoke: QRectF sceneRect()
    C._ZNK14QGraphicsScene9sceneRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "sceneRect", args)
  }

}

  // proto:  QGraphicsWidget * QGraphicsScene::activeWindow();
func (this *QGraphicsScene) activeWindow(args ...interface{}) () {
  // activeWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene12activeWindowEv
    // invoke: QGraphicsWidget * activeWindow()
    C._ZNK14QGraphicsScene12activeWindowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "activeWindow", args)
  }

}

  // proto:  QBrush QGraphicsScene::backgroundBrush();
func (this *QGraphicsScene) backgroundBrush(args ...interface{}) () {
  // backgroundBrush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene15backgroundBrushEv
    // invoke: QBrush backgroundBrush()
    C._ZNK14QGraphicsScene15backgroundBrushEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "backgroundBrush", args)
  }

}

  // proto:  QGraphicsItem * QGraphicsScene::itemAt(qreal x, qreal y, const QTransform & deviceTransform);
func (this *QGraphicsScene) itemAt(args ...interface{}) () {
  // itemAt(qreal, qreal, const class QTransform &)
  // itemAt(const class QPointF &, const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1][1] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene6itemAtEddRK10QTransform
    // invoke: QGraphicsItem * itemAt(qreal, qreal, const class QTransform &)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QTransform).qclsinst
    if false {fmt.Println(arg2)}
    C.demth_ZNK14QGraphicsScene6itemAtEddRK10QTransform(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZNK14QGraphicsScene6itemAtERK7QPointFRK10QTransform
    // invoke: QGraphicsItem * itemAt(const class QPointF &, const class QTransform &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTransform).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK14QGraphicsScene6itemAtERK7QPointFRK10QTransform(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "itemAt", args)
  }

}

  // proto:  void QGraphicsScene::advance();
func (this *QGraphicsScene) advance(args ...interface{}) () {
  // advance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7advanceEv
    // invoke: void advance()
    C._ZN14QGraphicsScene7advanceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "advance", args)
  }

}

  // proto:  void QGraphicsScene::setStickyFocus(bool enabled);
func (this *QGraphicsScene) setStickyFocus(args ...interface{}) () {
  // setStickyFocus(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene14setStickyFocusEb
    // invoke: void setStickyFocus(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScene14setStickyFocusEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setStickyFocus", args)
  }

}

  // proto:  QList<QGraphicsItem *> QGraphicsScene::selectedItems();
func (this *QGraphicsScene) selectedItems(args ...interface{}) () {
  // selectedItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene13selectedItemsEv
    // invoke: QList<QGraphicsItem *> selectedItems()
    C._ZNK14QGraphicsScene13selectedItemsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "selectedItems", args)
  }

}

  // proto:  void QGraphicsScene::clear();
func (this *QGraphicsScene) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene5clearEv
    // invoke: void clear()
    C._ZN14QGraphicsScene5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "clear", args)
  }

}

  // proto:  void QGraphicsScene::setActivePanel(QGraphicsItem * item);
func (this *QGraphicsScene) setActivePanel(args ...interface{}) () {
  // setActivePanel(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene14setActivePanelEP13QGraphicsItem
    // invoke: void setActivePanel(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScene14setActivePanelEP13QGraphicsItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setActivePanel", args)
  }

}

  // proto:  QGraphicsPixmapItem * QGraphicsScene::addPixmap(const QPixmap & pixmap);
func (this *QGraphicsScene) addPixmap(args ...interface{}) () {
  // addPixmap(const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene9addPixmapERK7QPixmap
    // invoke: QGraphicsPixmapItem * addPixmap(const class QPixmap &)
    var arg0 = args[0].(QPixmap).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScene9addPixmapERK7QPixmap(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addPixmap", args)
  }

}

  // proto:  QBrush QGraphicsScene::foregroundBrush();
func (this *QGraphicsScene) foregroundBrush(args ...interface{}) () {
  // foregroundBrush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene15foregroundBrushEv
    // invoke: QBrush foregroundBrush()
    C._ZNK14QGraphicsScene15foregroundBrushEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "foregroundBrush", args)
  }

}

  // proto:  QList<QGraphicsView *> QGraphicsScene::views();
func (this *QGraphicsScene) views(args ...interface{}) () {
  // views()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene5viewsEv
    // invoke: QList<QGraphicsView *> views()
    C._ZNK14QGraphicsScene5viewsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "views", args)
  }

}

  // proto:  void QGraphicsScene::~QGraphicsScene();
func (this *QGraphicsScene) FreeQGraphicsScene(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsScene", "~QGraphicsScene", args)
  }

}

  // proto:  QGraphicsRectItem * QGraphicsScene::addRect(qreal x, qreal y, qreal w, qreal h, const QPen & pen, const QBrush & brush);
func (this *QGraphicsScene) addRect(args ...interface{}) () {
  // addRect(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
  // addRect(const class QRectF &, const class QPen &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[0][5] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[1][2] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addRectEddddRK4QPenRK6QBrush
    // invoke: QGraphicsRectItem * addRect(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QPen).qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(QBrush).qclsinst
    if false {fmt.Println(arg5)}
    C.demth_ZN14QGraphicsScene7addRectEddddRK4QPenRK6QBrush(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  case 1:
    // invoke: _ZN14QGraphicsScene7addRectERK6QRectFRK4QPenRK6QBrush
    // invoke: QGraphicsRectItem * addRect(const class QRectF &, const class QPen &, const class QBrush &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPen).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QBrush).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN14QGraphicsScene7addRectERK6QRectFRK4QPenRK6QBrush(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addRect", args)
  }

}

  // proto:  int QGraphicsScene::bspTreeDepth();
func (this *QGraphicsScene) bspTreeDepth(args ...interface{}) () {
  // bspTreeDepth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene12bspTreeDepthEv
    // invoke: int bspTreeDepth()
    C._ZNK14QGraphicsScene12bspTreeDepthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "bspTreeDepth", args)
  }

}

  // proto:  void QGraphicsScene::setStyle(QStyle * style);
func (this *QGraphicsScene) setStyle(args ...interface{}) () {
  // setStyle(class QStyle *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyle{}) // "QStyle *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene8setStyleEP6QStyle
    // invoke: void setStyle(class QStyle *)
    var arg0 = args[0].(QStyle).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScene8setStyleEP6QStyle(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setStyle", args)
  }

}

  // proto:  void QGraphicsScene::setPalette(const QPalette & palette);
func (this *QGraphicsScene) setPalette(args ...interface{}) () {
  // setPalette(const class QPalette &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPalette{}) // "const QPalette &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10setPaletteERK8QPalette
    // invoke: void setPalette(const class QPalette &)
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScene10setPaletteERK8QPalette(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setPalette", args)
  }

}

  // proto:  void QGraphicsScene::setMinimumRenderSize(qreal minSize);
func (this *QGraphicsScene) setMinimumRenderSize(args ...interface{}) () {
  // setMinimumRenderSize(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene20setMinimumRenderSizeEd
    // invoke: void setMinimumRenderSize(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScene20setMinimumRenderSizeEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setMinimumRenderSize", args)
  }

}

  // proto:  QGraphicsItem * QGraphicsScene::mouseGrabberItem();
func (this *QGraphicsScene) mouseGrabberItem(args ...interface{}) () {
  // mouseGrabberItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene16mouseGrabberItemEv
    // invoke: QGraphicsItem * mouseGrabberItem()
    C._ZNK14QGraphicsScene16mouseGrabberItemEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "mouseGrabberItem", args)
  }

}

  // proto:  QGraphicsEllipseItem * QGraphicsScene::addEllipse(const QRectF & rect, const QPen & pen, const QBrush & brush);
func (this *QGraphicsScene) addEllipse(args ...interface{}) () {
  // addEllipse(const class QRectF &, const class QPen &, const class QBrush &)
  // addEllipse(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[0][2] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][4] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[1][5] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10addEllipseERK6QRectFRK4QPenRK6QBrush
    // invoke: QGraphicsEllipseItem * addEllipse(const class QRectF &, const class QPen &, const class QBrush &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPen).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QBrush).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN14QGraphicsScene10addEllipseERK6QRectFRK4QPenRK6QBrush(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN14QGraphicsScene10addEllipseEddddRK4QPenRK6QBrush
    // invoke: QGraphicsEllipseItem * addEllipse(qreal, qreal, qreal, qreal, const class QPen &, const class QBrush &)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QPen).qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(QBrush).qclsinst
    if false {fmt.Println(arg5)}
    C.demth_ZN14QGraphicsScene10addEllipseEddddRK4QPenRK6QBrush(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addEllipse", args)
  }

}

  // proto:  qreal QGraphicsScene::height();
func (this *QGraphicsScene) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene6heightEv
    // invoke: qreal height()
    C.demth_ZNK14QGraphicsScene6heightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "height", args)
  }

}

  // proto:  void QGraphicsScene::setSelectionArea(const QPainterPath & path, const QTransform & deviceTransform);
func (this *QGraphicsScene) setSelectionArea(args ...interface{}) () {
  // setSelectionArea(const class QPainterPath &, Qt::ItemSelectionMode, const class QTransform &)
  // setSelectionArea(const class QPainterPath &, const class QTransform &)
  // setSelectionArea(const class QPainterPath &, Qt::ItemSelectionOperation, Qt::ItemSelectionMode, const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[0][1] = qtrt.Int32Ty(false) // "Qt::ItemSelectionMode"
  vtys[0][2] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[1][1] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[2][1] = qtrt.Int32Ty(false) // "Qt::ItemSelectionOperation"
  vtys[2][2] = qtrt.Int32Ty(false) // "Qt::ItemSelectionMode"
  vtys[2][3] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathRK10QTransform
    // invoke: void setSelectionArea(const class QPainterPath &, const class QTransform &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTransform).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN14QGraphicsScene16setSelectionAreaERK12QPainterPathRK10QTransform(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setSelectionArea", args)
  }

}

  // proto:  QFont QGraphicsScene::font();
func (this *QGraphicsScene) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene4fontEv
    // invoke: QFont font()
    C._ZNK14QGraphicsScene4fontEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "font", args)
  }

}

  // proto:  void QGraphicsScene::clearSelection();
func (this *QGraphicsScene) clearSelection(args ...interface{}) () {
  // clearSelection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene14clearSelectionEv
    // invoke: void clearSelection()
    C._ZN14QGraphicsScene14clearSelectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "clearSelection", args)
  }

}

  // proto:  void QGraphicsScene::removeItem(QGraphicsItem * item);
func (this *QGraphicsScene) removeItem(args ...interface{}) () {
  // removeItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene10removeItemEP13QGraphicsItem
    // invoke: void removeItem(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScene10removeItemEP13QGraphicsItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "removeItem", args)
  }

}

  // proto:  void QGraphicsScene::setActiveWindow(QGraphicsWidget * widget);
func (this *QGraphicsScene) setActiveWindow(args ...interface{}) () {
  // setActiveWindow(class QGraphicsWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsWidget{}) // "QGraphicsWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene15setActiveWindowEP15QGraphicsWidget
    // invoke: void setActiveWindow(class QGraphicsWidget *)
    var arg0 = args[0].(QGraphicsWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScene15setActiveWindowEP15QGraphicsWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setActiveWindow", args)
  }

}

  // proto:  QGraphicsItem * QGraphicsScene::focusItem();
func (this *QGraphicsScene) focusItem(args ...interface{}) () {
  // focusItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene9focusItemEv
    // invoke: QGraphicsItem * focusItem()
    C._ZNK14QGraphicsScene9focusItemEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "focusItem", args)
  }

}

  // proto:  QGraphicsTextItem * QGraphicsScene::addText(const QString & text, const QFont & font);
func (this *QGraphicsScene) addText(args ...interface{}) () {
  // addText(const class QString &, const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addTextERK7QStringRK5QFont
    // invoke: QGraphicsTextItem * addText(const class QString &, const class QFont &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QFont).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN14QGraphicsScene7addTextERK7QStringRK5QFont(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addText", args)
  }

}

  // proto:  void QGraphicsScene::setSortCacheEnabled(bool enabled);
func (this *QGraphicsScene) setSortCacheEnabled(args ...interface{}) () {
  // setSortCacheEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene19setSortCacheEnabledEb
    // invoke: void setSortCacheEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScene19setSortCacheEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setSortCacheEnabled", args)
  }

}

  // proto:  void QGraphicsScene::destroyItemGroup(QGraphicsItemGroup * group);
func (this *QGraphicsScene) destroyItemGroup(args ...interface{}) () {
  // destroyItemGroup(class QGraphicsItemGroup *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItemGroup{}) // "QGraphicsItemGroup *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene16destroyItemGroupEP18QGraphicsItemGroup
    // invoke: void destroyItemGroup(class QGraphicsItemGroup *)
    var arg0 = args[0].(QGraphicsItemGroup).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScene16destroyItemGroupEP18QGraphicsItemGroup(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "destroyItemGroup", args)
  }

}

  // proto:  qreal QGraphicsScene::width();
func (this *QGraphicsScene) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene5widthEv
    // invoke: qreal width()
    C.demth_ZNK14QGraphicsScene5widthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "width", args)
  }

}

  // proto:  void QGraphicsScene::addItem(QGraphicsItem * item);
func (this *QGraphicsScene) addItem(args ...interface{}) () {
  // addItem(class QGraphicsItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsItem{}) // "QGraphicsItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addItemEP13QGraphicsItem
    // invoke: void addItem(class QGraphicsItem *)
    var arg0 = args[0].(QGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScene7addItemEP13QGraphicsItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addItem", args)
  }

}

  // proto:  void QGraphicsScene::setBackgroundBrush(const QBrush & brush);
func (this *QGraphicsScene) setBackgroundBrush(args ...interface{}) () {
  // setBackgroundBrush(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene18setBackgroundBrushERK6QBrush
    // invoke: void setBackgroundBrush(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScene18setBackgroundBrushERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setBackgroundBrush", args)
  }

}

  // proto:  QGraphicsItem * QGraphicsScene::activePanel();
func (this *QGraphicsScene) activePanel(args ...interface{}) () {
  // activePanel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene11activePanelEv
    // invoke: QGraphicsItem * activePanel()
    C._ZNK14QGraphicsScene11activePanelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "activePanel", args)
  }

}

  // proto:  QStyle * QGraphicsScene::style();
func (this *QGraphicsScene) style(args ...interface{}) () {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene5styleEv
    // invoke: QStyle * style()
    C._ZNK14QGraphicsScene5styleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "style", args)
  }

}

  // proto:  void QGraphicsScene::setFont(const QFont & font);
func (this *QGraphicsScene) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QGraphicsScene7setFontERK5QFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "setFont", args)
  }

}

  // proto:  QGraphicsPathItem * QGraphicsScene::addPath(const QPainterPath & path, const QPen & pen, const QBrush & brush);
func (this *QGraphicsScene) addPath(args ...interface{}) () {
  // addPath(const class QPainterPath &, const class QPen &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[0][1] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[0][2] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QGraphicsScene7addPathERK12QPainterPathRK4QPenRK6QBrush
    // invoke: QGraphicsPathItem * addPath(const class QPainterPath &, const class QPen &, const class QBrush &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPen).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QBrush).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN14QGraphicsScene7addPathERK12QPainterPathRK4QPenRK6QBrush(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "addPath", args)
  }

}

  // proto:  bool QGraphicsScene::stickyFocus();
func (this *QGraphicsScene) stickyFocus(args ...interface{}) () {
  // stickyFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QGraphicsScene11stickyFocusEv
    // invoke: bool stickyFocus()
    C._ZNK14QGraphicsScene11stickyFocusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsScene", "stickyFocus", args)
  }

}

// <= body block end
