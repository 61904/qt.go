//  header block begin
// /usr/include/qt/QtWidgets/qtabwidget.h
// #include <qtabwidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 81
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QTabWidget struct {
	*QWidget
}

func (this *QTabWidget) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}

// /usr/include/qt/QtWidgets/qtabwidget.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTabWidget) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:71
// index:0
// void QTabWidget(class QWidget *)
func NewQTabWidget(parent unsafe.Pointer) *QTabWidget {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidgetC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQTabWidgetFromPointer(cthis)
	return gothis
}
func NewQTabWidgetFromPointer(cthis unsafe.Pointer) *QTabWidget {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QTabWidget{bcthis0}
}

// /usr/include/qt/QtWidgets/qtabwidget.h:72
// index:0
// virtual
// void ~QTabWidget()
func DeleteQTabWidget(*QTabWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:74
// index:0
// int addTab(class QWidget *, const class QString &)
func (this *QTabWidget) AddTab(widget unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (, widget QWidget *, const QString & arg1), (widget, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget6addTabEP7QWidgetRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:75
// index:1
// int addTab(class QWidget *, const class QIcon &, const class QString &)
func (this *QTabWidget) AddTab_1(widget unsafe.Pointer, icon unsafe.Pointer, label unsafe.Pointer) {
	// 1: (, widget QWidget *, icon const QIcon &, label const QString &), (widget, icon, label)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget6addTabEP7QWidgetRK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget, icon, label)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:77
// index:0
// int insertTab(int, class QWidget *, const class QString &)
func (this *QTabWidget) InsertTab(index int, widget unsafe.Pointer, arg2 unsafe.Pointer) {
	// 0: (, index int, widget QWidget *, const QString & arg2), (&index, widget, arg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget9insertTabEiP7QWidgetRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, widget, arg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:78
// index:1
// int insertTab(int, class QWidget *, const class QIcon &, const class QString &)
func (this *QTabWidget) InsertTab_1(index int, widget unsafe.Pointer, icon unsafe.Pointer, label unsafe.Pointer) {
	// 1: (, index int, widget QWidget *, icon const QIcon &, label const QString &), (&index, widget, icon, label)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget9insertTabEiP7QWidgetRK5QIconRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, widget, icon, label)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:80
// index:0
// void removeTab(int)
func (this *QTabWidget) RemoveTab(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget9removeTabEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:82
// index:0
// bool isTabEnabled(int)
func (this *QTabWidget) IsTabEnabled(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget12isTabEnabledEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:83
// index:0
// void setTabEnabled(int, _Bool)
func (this *QTabWidget) SetTabEnabled(index int, arg1 bool) {
	// 0: (, index int, bool arg1), (&index, &arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget13setTabEnabledEib", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:85
// index:0
// QString tabText(int)
func (this *QTabWidget) TabText(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget7tabTextEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:86
// index:0
// void setTabText(int, const class QString &)
func (this *QTabWidget) SetTabText(index int, arg1 unsafe.Pointer) {
	// 0: (, index int, const QString & arg1), (&index, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget10setTabTextEiRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:88
// index:0
// QIcon tabIcon(int)
func (this *QTabWidget) TabIcon(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget7tabIconEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:89
// index:0
// void setTabIcon(int, const class QIcon &)
func (this *QTabWidget) SetTabIcon(index int, icon unsafe.Pointer) {
	// 0: (, index int, icon const QIcon &), (&index, icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget10setTabIconEiRK5QIcon", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:92
// index:0
// void setTabToolTip(int, const class QString &)
func (this *QTabWidget) SetTabToolTip(index int, tip unsafe.Pointer) {
	// 0: (, index int, tip const QString &), (&index, tip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget13setTabToolTipEiRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, tip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:93
// index:0
// QString tabToolTip(int)
func (this *QTabWidget) TabToolTip(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget10tabToolTipEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:97
// index:0
// void setTabWhatsThis(int, const class QString &)
func (this *QTabWidget) SetTabWhatsThis(index int, text unsafe.Pointer) {
	// 0: (, index int, text const QString &), (&index, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget15setTabWhatsThisEiRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:98
// index:0
// QString tabWhatsThis(int)
func (this *QTabWidget) TabWhatsThis(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget12tabWhatsThisEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:101
// index:0
// int currentIndex()
func (this *QTabWidget) CurrentIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget12currentIndexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:102
// index:0
// QWidget * currentWidget()
func (this *QTabWidget) CurrentWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget13currentWidgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:103
// index:0
// QWidget * widget(int)
func (this *QTabWidget) Widget(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget6widgetEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:104
// index:0
// int indexOf(class QWidget *)
func (this *QTabWidget) IndexOf(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget7indexOfEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:105
// index:0
// int count()
func (this *QTabWidget) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget5countEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:109
// index:0
// QTabWidget::TabPosition tabPosition()
func (this *QTabWidget) TabPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget11tabPositionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:110
// index:0
// void setTabPosition(enum QTabWidget::TabPosition)
func (this *QTabWidget) SetTabPosition(arg0 int) {
	// 0: (, QTabWidget::TabPosition arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget14setTabPositionENS_11TabPositionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:112
// index:0
// bool tabsClosable()
func (this *QTabWidget) TabsClosable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget12tabsClosableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:113
// index:0
// void setTabsClosable(_Bool)
func (this *QTabWidget) SetTabsClosable(closeable bool) {
	// 0: (, closeable bool), (&closeable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget15setTabsClosableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &closeable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:115
// index:0
// bool isMovable()
func (this *QTabWidget) IsMovable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget9isMovableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:116
// index:0
// void setMovable(_Bool)
func (this *QTabWidget) SetMovable(movable bool) {
	// 0: (, movable bool), (&movable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget10setMovableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &movable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:120
// index:0
// QTabWidget::TabShape tabShape()
func (this *QTabWidget) TabShape() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget8tabShapeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:121
// index:0
// void setTabShape(enum QTabWidget::TabShape)
func (this *QTabWidget) SetTabShape(s int) {
	// 0: (, s QTabWidget::TabShape), (&s)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget11setTabShapeENS_8TabShapeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &s)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:123
// index:0
// virtual
// QSize sizeHint()
func (this *QTabWidget) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:124
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QTabWidget) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:125
// index:0
// virtual
// int heightForWidth(int)
func (this *QTabWidget) HeightForWidth(width int) {
	// 0: (, width int), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget14heightForWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:126
// index:0
// virtual
// bool hasHeightForWidth()
func (this *QTabWidget) HasHeightForWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget17hasHeightForWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:128
// index:0
// void setCornerWidget(class QWidget *, Qt::Corner)
func (this *QTabWidget) SetCornerWidget(w unsafe.Pointer, corner int) {
	// 0: (, w QWidget *, corner Qt::Corner), (w, &corner)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget15setCornerWidgetEP7QWidgetN2Qt6CornerE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), w, &corner)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:129
// index:0
// QWidget * cornerWidget(Qt::Corner)
func (this *QTabWidget) CornerWidget(corner int) {
	// 0: (, corner Qt::Corner), (&corner)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget12cornerWidgetEN2Qt6CornerE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &corner)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:131
// index:0
// Qt::TextElideMode elideMode()
func (this *QTabWidget) ElideMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget9elideModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:132
// index:0
// void setElideMode(Qt::TextElideMode)
func (this *QTabWidget) SetElideMode(arg0 int) {
	// 0: (, Qt::TextElideMode arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget12setElideModeEN2Qt13TextElideModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:134
// index:0
// QSize iconSize()
func (this *QTabWidget) IconSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget8iconSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:135
// index:0
// void setIconSize(const class QSize &)
func (this *QTabWidget) SetIconSize(size unsafe.Pointer) {
	// 0: (, size const QSize &), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget11setIconSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:137
// index:0
// bool usesScrollButtons()
func (this *QTabWidget) UsesScrollButtons() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget17usesScrollButtonsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:138
// index:0
// void setUsesScrollButtons(_Bool)
func (this *QTabWidget) SetUsesScrollButtons(useButtons bool) {
	// 0: (, useButtons bool), (&useButtons)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget20setUsesScrollButtonsEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &useButtons)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:140
// index:0
// bool documentMode()
func (this *QTabWidget) DocumentMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget12documentModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:141
// index:0
// void setDocumentMode(_Bool)
func (this *QTabWidget) SetDocumentMode(set bool) {
	// 0: (, set bool), (&set)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget15setDocumentModeEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &set)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:143
// index:0
// bool tabBarAutoHide()
func (this *QTabWidget) TabBarAutoHide() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget14tabBarAutoHideEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:144
// index:0
// void setTabBarAutoHide(_Bool)
func (this *QTabWidget) SetTabBarAutoHide(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget17setTabBarAutoHideEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:146
// index:0
// void clear()
func (this *QTabWidget) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:148
// index:0
// QTabBar * tabBar()
func (this *QTabWidget) TabBar() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget6tabBarEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:151
// index:0
// void setCurrentIndex(int)
func (this *QTabWidget) SetCurrentIndex(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget15setCurrentIndexEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:152
// index:0
// void setCurrentWidget(class QWidget *)
func (this *QTabWidget) SetCurrentWidget(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget16setCurrentWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:155
// index:0
// void currentChanged(int)
func (this *QTabWidget) CurrentChanged(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget14currentChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:156
// index:0
// void tabCloseRequested(int)
func (this *QTabWidget) TabCloseRequested(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget17tabCloseRequestedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:157
// index:0
// void tabBarClicked(int)
func (this *QTabWidget) TabBarClicked(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget13tabBarClickedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:158
// index:0
// void tabBarDoubleClicked(int)
func (this *QTabWidget) TabBarDoubleClicked(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget19tabBarDoubleClickedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:161
// index:0
// virtual
// void tabInserted(int)
func (this *QTabWidget) TabInserted(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget11tabInsertedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:162
// index:0
// virtual
// void tabRemoved(int)
func (this *QTabWidget) TabRemoved(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget10tabRemovedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:164
// index:0
// virtual
// void showEvent(class QShowEvent *)
func (this *QTabWidget) ShowEvent(arg0 unsafe.Pointer) {
	// 0: (, QShowEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget9showEventEP10QShowEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:165
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QTabWidget) ResizeEvent(arg0 unsafe.Pointer) {
	// 0: (, QResizeEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:166
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QTabWidget) KeyPressEvent(arg0 unsafe.Pointer) {
	// 0: (, QKeyEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:167
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QTabWidget) PaintEvent(arg0 unsafe.Pointer) {
	// 0: (, QPaintEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:168
// index:0
// void setTabBar(class QTabBar *)
func (this *QTabWidget) SetTabBar(arg0 unsafe.Pointer) {
	// 0: (, QTabBar * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget9setTabBarEP7QTabBar", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:169
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QTabWidget) ChangeEvent(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:170
// index:0
// virtual
// bool event(class QEvent *)
func (this *QTabWidget) Event(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:171
// index:0
// void initStyleOption(class QStyleOptionTabWidgetFrame *)
func (this *QTabWidget) InitStyleOption(option unsafe.Pointer) {
	// 0: (, option QStyleOptionTabWidgetFrame *), (option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget15initStyleOptionEP26QStyleOptionTabWidgetFrame", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end