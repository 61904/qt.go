package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtWidgets/qabstractbutton.h
// dst-file: /src/widgets/qabstractbutton.go
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
  // proto:  QSize QAbstractButton::iconSize();
extern void _ZNK15QAbstractButton8iconSizeEv(void* qthis);
  // proto:  void QAbstractButton::click();
extern void _ZN15QAbstractButton5clickEv(void* qthis);
  // proto:  void QAbstractButton::~QAbstractButton();
extern void _ZN15QAbstractButtonD0Ev(void* qthis);
  // proto:  void QAbstractButton::setChecked(bool );
extern void _ZN15QAbstractButton10setCheckedEb(void* qthis, bool arg0);
  // proto:  QKeySequence QAbstractButton::shortcut();
extern void _ZNK15QAbstractButton8shortcutEv(void* qthis);
  // proto:  QButtonGroup * QAbstractButton::group();
extern void _ZNK15QAbstractButton5groupEv(void* qthis);
  // proto:  bool QAbstractButton::isCheckable();
extern void _ZNK15QAbstractButton11isCheckableEv(void* qthis);
  // proto:  void QAbstractButton::QAbstractButton(QWidget * parent);
extern void* dector_ZN15QAbstractButtonC1EP7QWidget(void* arg0);
extern void _ZN15QAbstractButtonC1EP7QWidget(void* qthis, void* arg0);
  // proto:  bool QAbstractButton::isDown();
extern void _ZNK15QAbstractButton6isDownEv(void* qthis);
  // proto:  void QAbstractButton::setAutoExclusive(bool );
extern void _ZN15QAbstractButton16setAutoExclusiveEb(void* qthis, bool arg0);
  // proto:  const QMetaObject * QAbstractButton::metaObject();
extern void _ZNK15QAbstractButton10metaObjectEv(void* qthis);
  // proto:  bool QAbstractButton::isChecked();
extern void _ZNK15QAbstractButton9isCheckedEv(void* qthis);
  // proto:  void QAbstractButton::setAutoRepeatDelay(int );
extern void _ZN15QAbstractButton18setAutoRepeatDelayEi(void* qthis, int32_t arg0);
  // proto:  int QAbstractButton::autoRepeatDelay();
extern void _ZNK15QAbstractButton15autoRepeatDelayEv(void* qthis);
  // proto:  bool QAbstractButton::autoExclusive();
extern void _ZNK15QAbstractButton13autoExclusiveEv(void* qthis);
  // proto:  void QAbstractButton::toggle();
extern void _ZN15QAbstractButton6toggleEv(void* qthis);
  // proto:  void QAbstractButton::setIcon(const QIcon & icon);
extern void _ZN15QAbstractButton7setIconERK5QIcon(void* qthis, void* arg0);
  // proto:  void QAbstractButton::setAutoRepeatInterval(int );
extern void _ZN15QAbstractButton21setAutoRepeatIntervalEi(void* qthis, int32_t arg0);
  // proto:  void QAbstractButton::setAutoRepeat(bool );
extern void _ZN15QAbstractButton13setAutoRepeatEb(void* qthis, bool arg0);
  // proto:  void QAbstractButton::QAbstractButton(const QAbstractButton & );
extern void* dector_ZN15QAbstractButtonC1ERKS_(void* arg0);
extern void _ZN15QAbstractButtonC1ERKS_(void* qthis, void* arg0);
  // proto:  void QAbstractButton::animateClick(int msec);
extern void _ZN15QAbstractButton12animateClickEi(void* qthis, int32_t arg0);
  // proto:  void QAbstractButton::setDown(bool );
extern void _ZN15QAbstractButton7setDownEb(void* qthis, bool arg0);
  // proto:  QString QAbstractButton::text();
extern void _ZNK15QAbstractButton4textEv(void* qthis);
  // proto:  void QAbstractButton::setShortcut(const QKeySequence & key);
extern void _ZN15QAbstractButton11setShortcutERK12QKeySequence(void* qthis, void* arg0);
  // proto:  void QAbstractButton::setCheckable(bool );
extern void _ZN15QAbstractButton12setCheckableEb(void* qthis, bool arg0);
  // proto:  QIcon QAbstractButton::icon();
extern void _ZNK15QAbstractButton4iconEv(void* qthis);
  // proto:  void QAbstractButton::setText(const QString & text);
extern void _ZN15QAbstractButton7setTextERK7QString(void* qthis, void* arg0);
  // proto:  int QAbstractButton::autoRepeatInterval();
extern void _ZNK15QAbstractButton18autoRepeatIntervalEv(void* qthis);
  // proto:  bool QAbstractButton::autoRepeat();
extern void _ZNK15QAbstractButton10autoRepeatEv(void* qthis);
  // proto:  void QAbstractButton::setIconSize(const QSize & size);
extern void _ZN15QAbstractButton11setIconSizeERK5QSize(void* qthis, void* arg0);
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

// class sizeof(QAbstractButton)=1
type QAbstractButton struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
//  _released QAbstractButton_released_signal;
//  _clicked QAbstractButton_clicked_signal;
//  _pressed QAbstractButton_pressed_signal;
//  _toggled QAbstractButton_toggled_signal;
}

  // proto:  QSize QAbstractButton::iconSize();
func (this *QAbstractButton) iconSize(args ...interface{}) () {
  // iconSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton8iconSizeEv
    // invoke: QSize iconSize()
    C._ZNK15QAbstractButton8iconSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "iconSize", args)
  }

}

  // proto:  void QAbstractButton::click();
func (this *QAbstractButton) click(args ...interface{}) () {
  // click()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton5clickEv
    // invoke: void click()
    C._ZN15QAbstractButton5clickEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "click", args)
  }

}

  // proto:  void QAbstractButton::~QAbstractButton();
func (this *QAbstractButton) FreeQAbstractButton(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractButton", "~QAbstractButton", args)
  }

}

  // proto:  void QAbstractButton::setChecked(bool );
func (this *QAbstractButton) setChecked(args ...interface{}) () {
  // setChecked(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton10setCheckedEb
    // invoke: void setChecked(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN15QAbstractButton10setCheckedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setChecked", args)
  }

}

  // proto:  QKeySequence QAbstractButton::shortcut();
func (this *QAbstractButton) shortcut(args ...interface{}) () {
  // shortcut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton8shortcutEv
    // invoke: QKeySequence shortcut()
    C._ZNK15QAbstractButton8shortcutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "shortcut", args)
  }

}

  // proto:  QButtonGroup * QAbstractButton::group();
func (this *QAbstractButton) group(args ...interface{}) () {
  // group()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton5groupEv
    // invoke: QButtonGroup * group()
    C._ZNK15QAbstractButton5groupEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "group", args)
  }

}

  // proto:  bool QAbstractButton::isCheckable();
func (this *QAbstractButton) isCheckable(args ...interface{}) () {
  // isCheckable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton11isCheckableEv
    // invoke: bool isCheckable()
    C._ZNK15QAbstractButton11isCheckableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "isCheckable", args)
  }

}

  // proto:  void QAbstractButton::QAbstractButton(QWidget * parent);
func NewQAbstractButton(args ...interface{}) QAbstractButton {
  return QAbstractButton{}
}

  // proto:  bool QAbstractButton::isDown();
func (this *QAbstractButton) isDown(args ...interface{}) () {
  // isDown()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton6isDownEv
    // invoke: bool isDown()
    C._ZNK15QAbstractButton6isDownEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "isDown", args)
  }

}

  // proto:  void QAbstractButton::setAutoExclusive(bool );
func (this *QAbstractButton) setAutoExclusive(args ...interface{}) () {
  // setAutoExclusive(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton16setAutoExclusiveEb
    // invoke: void setAutoExclusive(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN15QAbstractButton16setAutoExclusiveEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setAutoExclusive", args)
  }

}

  // proto:  const QMetaObject * QAbstractButton::metaObject();
func (this *QAbstractButton) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK15QAbstractButton10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "metaObject", args)
  }

}

  // proto:  bool QAbstractButton::isChecked();
func (this *QAbstractButton) isChecked(args ...interface{}) () {
  // isChecked()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton9isCheckedEv
    // invoke: bool isChecked()
    C._ZNK15QAbstractButton9isCheckedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "isChecked", args)
  }

}

  // proto:  void QAbstractButton::setAutoRepeatDelay(int );
func (this *QAbstractButton) setAutoRepeatDelay(args ...interface{}) () {
  // setAutoRepeatDelay(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton18setAutoRepeatDelayEi
    // invoke: void setAutoRepeatDelay(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN15QAbstractButton18setAutoRepeatDelayEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setAutoRepeatDelay", args)
  }

}

  // proto:  int QAbstractButton::autoRepeatDelay();
func (this *QAbstractButton) autoRepeatDelay(args ...interface{}) () {
  // autoRepeatDelay()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton15autoRepeatDelayEv
    // invoke: int autoRepeatDelay()
    C._ZNK15QAbstractButton15autoRepeatDelayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "autoRepeatDelay", args)
  }

}

  // proto:  bool QAbstractButton::autoExclusive();
func (this *QAbstractButton) autoExclusive(args ...interface{}) () {
  // autoExclusive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton13autoExclusiveEv
    // invoke: bool autoExclusive()
    C._ZNK15QAbstractButton13autoExclusiveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "autoExclusive", args)
  }

}

  // proto:  void QAbstractButton::toggle();
func (this *QAbstractButton) toggle(args ...interface{}) () {
  // toggle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton6toggleEv
    // invoke: void toggle()
    C._ZN15QAbstractButton6toggleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "toggle", args)
  }

}

  // proto:  void QAbstractButton::setIcon(const QIcon & icon);
func (this *QAbstractButton) setIcon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton7setIconERK5QIcon
    // invoke: void setIcon(const class QIcon &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QAbstractButton7setIconERK5QIcon(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setIcon", args)
  }

}

  // proto:  void QAbstractButton::setAutoRepeatInterval(int );
func (this *QAbstractButton) setAutoRepeatInterval(args ...interface{}) () {
  // setAutoRepeatInterval(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton21setAutoRepeatIntervalEi
    // invoke: void setAutoRepeatInterval(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN15QAbstractButton21setAutoRepeatIntervalEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setAutoRepeatInterval", args)
  }

}

  // proto:  void QAbstractButton::setAutoRepeat(bool );
func (this *QAbstractButton) setAutoRepeat(args ...interface{}) () {
  // setAutoRepeat(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton13setAutoRepeatEb
    // invoke: void setAutoRepeat(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN15QAbstractButton13setAutoRepeatEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setAutoRepeat", args)
  }

}

  // proto:  void QAbstractButton::animateClick(int msec);
func (this *QAbstractButton) animateClick(args ...interface{}) () {
  // animateClick(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton12animateClickEi
    // invoke: void animateClick(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN15QAbstractButton12animateClickEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "animateClick", args)
  }

}

  // proto:  void QAbstractButton::setDown(bool );
func (this *QAbstractButton) setDown(args ...interface{}) () {
  // setDown(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton7setDownEb
    // invoke: void setDown(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN15QAbstractButton7setDownEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setDown", args)
  }

}

  // proto:  QString QAbstractButton::text();
func (this *QAbstractButton) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton4textEv
    // invoke: QString text()
    C._ZNK15QAbstractButton4textEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "text", args)
  }

}

  // proto:  void QAbstractButton::setShortcut(const QKeySequence & key);
func (this *QAbstractButton) setShortcut(args ...interface{}) () {
  // setShortcut(const class QKeySequence &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QKeySequence{}) // "const QKeySequence &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton11setShortcutERK12QKeySequence
    // invoke: void setShortcut(const class QKeySequence &)
    var arg0 = args[0].(QKeySequence).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QAbstractButton11setShortcutERK12QKeySequence(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setShortcut", args)
  }

}

  // proto:  void QAbstractButton::setCheckable(bool );
func (this *QAbstractButton) setCheckable(args ...interface{}) () {
  // setCheckable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton12setCheckableEb
    // invoke: void setCheckable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN15QAbstractButton12setCheckableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setCheckable", args)
  }

}

  // proto:  QIcon QAbstractButton::icon();
func (this *QAbstractButton) icon(args ...interface{}) () {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton4iconEv
    // invoke: QIcon icon()
    C._ZNK15QAbstractButton4iconEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "icon", args)
  }

}

  // proto:  void QAbstractButton::setText(const QString & text);
func (this *QAbstractButton) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QAbstractButton7setTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setText", args)
  }

}

  // proto:  int QAbstractButton::autoRepeatInterval();
func (this *QAbstractButton) autoRepeatInterval(args ...interface{}) () {
  // autoRepeatInterval()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton18autoRepeatIntervalEv
    // invoke: int autoRepeatInterval()
    C._ZNK15QAbstractButton18autoRepeatIntervalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "autoRepeatInterval", args)
  }

}

  // proto:  bool QAbstractButton::autoRepeat();
func (this *QAbstractButton) autoRepeat(args ...interface{}) () {
  // autoRepeat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractButton10autoRepeatEv
    // invoke: bool autoRepeat()
    C._ZNK15QAbstractButton10autoRepeatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractButton", "autoRepeat", args)
  }

}

  // proto:  void QAbstractButton::setIconSize(const QSize & size);
func (this *QAbstractButton) setIconSize(args ...interface{}) () {
  // setIconSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractButton11setIconSizeERK5QSize
    // invoke: void setIconSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QAbstractButton11setIconSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractButton", "setIconSize", args)
  }

}

// <= body block end

