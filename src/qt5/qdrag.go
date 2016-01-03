package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtGui/qdrag.h
// dst-file: /src/gui/qdrag.go
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
  // proto:  QObject * QDrag::target();
extern void _ZNK5QDrag6targetEv(void* qthis);
  // proto:  QMimeData * QDrag::mimeData();
extern void _ZNK5QDrag8mimeDataEv(void* qthis);
  // proto:  void QDrag::QDrag(QObject * dragSource);
extern void* dector_ZN5QDragC1EP7QObject(void* arg0);
extern void _ZN5QDragC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QDrag::~QDrag();
extern void _ZN5QDragD0Ev(void* qthis);
  // proto:  void QDrag::QDrag(const QDrag & );
extern void* dector_ZN5QDragC1ERKS_(void* arg0);
extern void _ZN5QDragC1ERKS_(void* qthis, void* arg0);
  // proto:  void QDrag::setHotSpot(const QPoint & hotspot);
extern void _ZN5QDrag10setHotSpotERK6QPoint(void* qthis, void* arg0);
  // proto:  const QMetaObject * QDrag::metaObject();
extern void _ZNK5QDrag10metaObjectEv(void* qthis);
  // proto:  void QDrag::setMimeData(QMimeData * data);
extern void _ZN5QDrag11setMimeDataEP9QMimeData(void* qthis, void* arg0);
  // proto:  QPixmap QDrag::pixmap();
extern void _ZNK5QDrag6pixmapEv(void* qthis);
  // proto:  QPoint QDrag::hotSpot();
extern void _ZNK5QDrag7hotSpotEv(void* qthis);
  // proto:  void QDrag::setPixmap(const QPixmap & );
extern void _ZN5QDrag9setPixmapERK7QPixmap(void* qthis, void* arg0);
  // proto:  QObject * QDrag::source();
extern void _ZNK5QDrag6sourceEv(void* qthis);
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

// class sizeof(QDrag)=1
type QDrag struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _targetChanged QDrag_targetChanged_signal;
//  _actionChanged QDrag_actionChanged_signal;
}

  // proto:  QObject * QDrag::target();
func (this *QDrag) target(args ...interface{}) () {
  // target()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDrag6targetEv
    // invoke: QObject * target()
    C._ZNK5QDrag6targetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "target", args)
  }

}

  // proto:  QMimeData * QDrag::mimeData();
func (this *QDrag) mimeData(args ...interface{}) () {
  // mimeData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDrag8mimeDataEv
    // invoke: QMimeData * mimeData()
    C._ZNK5QDrag8mimeDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "mimeData", args)
  }

}

  // proto:  void QDrag::QDrag(QObject * dragSource);
func NewQDrag(args ...interface{}) QDrag {
  return QDrag{}
}

  // proto:  void QDrag::~QDrag();
func (this *QDrag) FreeQDrag(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDrag", "~QDrag", args)
  }

}

  // proto:  void QDrag::setHotSpot(const QPoint & hotspot);
func (this *QDrag) setHotSpot(args ...interface{}) () {
  // setHotSpot(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDrag10setHotSpotERK6QPoint
    // invoke: void setHotSpot(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN5QDrag10setHotSpotERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDrag", "setHotSpot", args)
  }

}

  // proto:  const QMetaObject * QDrag::metaObject();
func (this *QDrag) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDrag10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK5QDrag10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "metaObject", args)
  }

}

  // proto:  void QDrag::setMimeData(QMimeData * data);
func (this *QDrag) setMimeData(args ...interface{}) () {
  // setMimeData(class QMimeData *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMimeData{}) // "QMimeData *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDrag11setMimeDataEP9QMimeData
    // invoke: void setMimeData(class QMimeData *)
    var arg0 = args[0].(QMimeData).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN5QDrag11setMimeDataEP9QMimeData(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDrag", "setMimeData", args)
  }

}

  // proto:  QPixmap QDrag::pixmap();
func (this *QDrag) pixmap(args ...interface{}) () {
  // pixmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDrag6pixmapEv
    // invoke: QPixmap pixmap()
    C._ZNK5QDrag6pixmapEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "pixmap", args)
  }

}

  // proto:  QPoint QDrag::hotSpot();
func (this *QDrag) hotSpot(args ...interface{}) () {
  // hotSpot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDrag7hotSpotEv
    // invoke: QPoint hotSpot()
    C._ZNK5QDrag7hotSpotEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "hotSpot", args)
  }

}

  // proto:  void QDrag::setPixmap(const QPixmap & );
func (this *QDrag) setPixmap(args ...interface{}) () {
  // setPixmap(const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDrag9setPixmapERK7QPixmap
    // invoke: void setPixmap(const class QPixmap &)
    var arg0 = args[0].(QPixmap).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN5QDrag9setPixmapERK7QPixmap(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDrag", "setPixmap", args)
  }

}

  // proto:  QObject * QDrag::source();
func (this *QDrag) source(args ...interface{}) () {
  // source()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDrag6sourceEv
    // invoke: QObject * source()
    C._ZNK5QDrag6sourceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "source", args)
  }

}

// <= body block end
