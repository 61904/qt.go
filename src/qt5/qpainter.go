package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtGui/qpainter.h
// dst-file: /src/gui/qpainter.go
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
  // proto:  QRectF QPainter::boundingRect(const QRectF & rect, const QString & text, const QTextOption & o);
extern void _ZN8QPainter12boundingRectERK6QRectFRK7QStringRK11QTextOption(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QPainter::drawPicture(const QPointF & p, const QPicture & picture);
extern void _ZN8QPainter11drawPictureERK7QPointFRK8QPicture(void* qthis, void* arg0, void* arg1);
  // proto:  const QMatrix & QPainter::worldMatrix();
extern void _ZNK8QPainter11worldMatrixEv(void* qthis);
  // proto:  void QPainter::drawText(const QPointF & p, const QString & str, int tf, int justificationPadding);
extern void _ZN8QPainter8drawTextERK7QPointFRK7QStringii(void* qthis, void* arg0, void* arg1, int32_t arg2, int32_t arg3);
  // proto:  void QPainter::fillRect(int x, int y, int w, int h, const QColor & color);
extern void demth_ZN8QPainter8fillRectEiiiiRK6QColor(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, void* arg4);
  // proto:  const QMatrix & QPainter::matrix();
extern void _ZNK8QPainter6matrixEv(void* qthis);
  // proto:  qreal QPainter::opacity();
extern void _ZNK8QPainter7opacityEv(void* qthis);
  // proto:  void QPainter::drawText(int x, int y, const QString & s);
extern void demth_ZN8QPainter8drawTextEiiRK7QString(void* qthis, int32_t arg0, int32_t arg1, void* arg2);
  // proto:  void QPainter::drawTiledPixmap(const QRectF & rect, const QPixmap & pm, const QPointF & offset);
extern void _ZN8QPainter15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QPainter::setBackground(const QBrush & bg);
extern void _ZN8QPainter13setBackgroundERK6QBrush(void* qthis, void* arg0);
  // proto:  QRect QPainter::boundingRect(const QRect & rect, int flags, const QString & text);
extern void _ZN8QPainter12boundingRectERK5QRectiRK7QString(void* qthis, void* arg0, int32_t arg1, void* arg2);
  // proto:  void QPainter::drawChord(const QRectF & rect, int a, int alen);
extern void _ZN8QPainter9drawChordERK6QRectFii(void* qthis, void* arg0, int32_t arg1, int32_t arg2);
  // proto:  void QPainter::drawImage(const QRectF & r, const QImage & image);
extern void demth_ZN8QPainter9drawImageERK6QRectFRK6QImage(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::setClipping(bool enable);
extern void _ZN8QPainter11setClippingEb(void* qthis, bool arg0);
  // proto:  void QPainter::setBrush(const QBrush & brush);
extern void _ZN8QPainter8setBrushERK6QBrush(void* qthis, void* arg0);
  // proto:  void QPainter::setMatrix(const QMatrix & matrix, bool combine);
extern void _ZN8QPainter9setMatrixERK7QMatrixb(void* qthis, void* arg0, bool arg1);
  // proto:  void QPainter::drawChord(const QRect & , int a, int alen);
extern void demth_ZN8QPainter9drawChordERK5QRectii(void* qthis, void* arg0, int32_t arg1, int32_t arg2);
  // proto:  void QPainter::eraseRect(const QRectF & );
extern void _ZN8QPainter9eraseRectERK6QRectF(void* qthis, void* arg0);
  // proto:  void QPainter::translate(const QPoint & offset);
extern void demth_ZN8QPainter9translateERK6QPoint(void* qthis, void* arg0);
  // proto:  bool QPainter::viewTransformEnabled();
extern void _ZNK8QPainter20viewTransformEnabledEv(void* qthis);
  // proto:  void QPainter::setPen(const QPen & pen);
extern void _ZN8QPainter6setPenERK4QPen(void* qthis, void* arg0);
  // proto:  void QPainter::drawLines(const QLineF * lines, int lineCount);
extern void _ZN8QPainter9drawLinesEPK6QLineFi(void* qthis, void* arg0, int32_t arg1);
  // proto:  void QPainter::setBrushOrigin(int x, int y);
extern void demth_ZN8QPainter14setBrushOriginEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  const QTransform & QPainter::worldTransform();
extern void _ZNK8QPainter14worldTransformEv(void* qthis);
  // proto:  void QPainter::drawRects(const QRect * rects, int rectCount);
extern void _ZN8QPainter9drawRectsEPK5QRecti(void* qthis, void* arg0, int32_t arg1);
  // proto:  void QPainter::drawEllipse(const QPoint & center, int rx, int ry);
extern void demth_ZN8QPainter11drawEllipseERK6QPointii(void* qthis, void* arg0, int32_t arg1, int32_t arg2);
  // proto:  void QPainter::drawArc(int x, int y, int w, int h, int a, int alen);
extern void demth_ZN8QPainter7drawArcEiiiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4, int32_t arg5);
  // proto:  void QPainter::drawPolyline(const QPolygonF & polyline);
extern void demth_ZN8QPainter12drawPolylineERK9QPolygonF(void* qthis, void* arg0);
  // proto:  bool QPainter::hasClipping();
extern void _ZNK8QPainter11hasClippingEv(void* qthis);
  // proto:  void QPainter::drawPixmap(const QRectF & targetRect, const QPixmap & pixmap, const QRectF & sourceRect);
extern void _ZN8QPainter10drawPixmapERK6QRectFRK7QPixmapS2_(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QPainter::drawStaticText(int left, int top, const QStaticText & staticText);
extern void demth_ZN8QPainter14drawStaticTextEiiRK11QStaticText(void* qthis, int32_t arg0, int32_t arg1, void* arg2);
  // proto:  void QPainter::strokePath(const QPainterPath & path, const QPen & pen);
extern void _ZN8QPainter10strokePathERK12QPainterPathRK4QPen(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::drawPixmap(int x, int y, const QPixmap & pm, int sx, int sy, int sw, int sh);
extern void demth_ZN8QPainter10drawPixmapEiiRK7QPixmapiiii(void* qthis, int32_t arg0, int32_t arg1, void* arg2, int32_t arg3, int32_t arg4, int32_t arg5, int32_t arg6);
  // proto:  void QPainter::drawRects(const QRectF * rects, int rectCount);
extern void _ZN8QPainter9drawRectsEPK6QRectFi(void* qthis, void* arg0, int32_t arg1);
  // proto:  void QPainter::drawConvexPolygon(const QPoint * points, int pointCount);
extern void _ZN8QPainter17drawConvexPolygonEPK6QPointi(void* qthis, void* arg0, int32_t arg1);
  // proto:  void QPainter::drawPath(const QPainterPath & path);
extern void _ZN8QPainter8drawPathERK12QPainterPath(void* qthis, void* arg0);
  // proto:  void QPainter::drawPixmap(int x, int y, const QPixmap & pm);
extern void demth_ZN8QPainter10drawPixmapEiiRK7QPixmap(void* qthis, int32_t arg0, int32_t arg1, void* arg2);
  // proto:  QMatrix QPainter::combinedMatrix();
extern void _ZNK8QPainter14combinedMatrixEv(void* qthis);
  // proto:  void QPainter::setMatrixEnabled(bool enabled);
extern void _ZN8QPainter16setMatrixEnabledEb(void* qthis, bool arg0);
  // proto:  void QPainter::drawPolyline(const QPolygon & polygon);
extern void demth_ZN8QPainter12drawPolylineERK8QPolygon(void* qthis, void* arg0);
  // proto:  void QPainter::drawTiledPixmap(const QRect & , const QPixmap & , const QPoint & );
extern void demth_ZN8QPainter15drawTiledPixmapERK5QRectRK7QPixmapRK6QPoint(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QPainter::setFont(const QFont & f);
extern void _ZN8QPainter7setFontERK5QFont(void* qthis, void* arg0);
  // proto:  void QPainter::drawChord(int x, int y, int w, int h, int a, int alen);
extern void demth_ZN8QPainter9drawChordEiiiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4, int32_t arg5);
  // proto:  void QPainter::drawPixmap(int x, int y, int w, int h, const QPixmap & pm);
extern void demth_ZN8QPainter10drawPixmapEiiiiRK7QPixmap(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, void* arg4);
  // proto:  void QPainter::setWindow(const QRect & window);
extern void _ZN8QPainter9setWindowERK5QRect(void* qthis, void* arg0);
  // proto:  const QMatrix & QPainter::deviceMatrix();
extern void _ZNK8QPainter12deviceMatrixEv(void* qthis);
  // proto:  void QPainter::drawLines(const QPointF * pointPairs, int lineCount);
extern void _ZN8QPainter9drawLinesEPK7QPointFi(void* qthis, void* arg0, int32_t arg1);
  // proto:  void QPainter::drawPixmap(const QPointF & p, const QPixmap & pm);
extern void _ZN8QPainter10drawPixmapERK7QPointFRK7QPixmap(void* qthis, void* arg0, void* arg1);
  // proto:  QRect QPainter::boundingRect(int x, int y, int w, int h, int flags, const QString & text);
extern void demth_ZN8QPainter12boundingRectEiiiiiRK7QString(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4, void* arg5);
  // proto:  void QPainter::drawLines(const QLine * lines, int lineCount);
extern void _ZN8QPainter9drawLinesEPK5QLinei(void* qthis, void* arg0, int32_t arg1);
  // proto:  void QPainter::drawPie(int x, int y, int w, int h, int a, int alen);
extern void demth_ZN8QPainter7drawPieEiiiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4, int32_t arg5);
  // proto:  void QPainter::drawPixmap(const QPoint & p, const QPixmap & pm, const QRect & sr);
extern void demth_ZN8QPainter10drawPixmapERK6QPointRK7QPixmapRK5QRect(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QPainter::drawStaticText(const QPointF & topLeftPosition, const QStaticText & staticText);
extern void _ZN8QPainter14drawStaticTextERK7QPointFRK11QStaticText(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::setWorldMatrixEnabled(bool enabled);
extern void _ZN8QPainter21setWorldMatrixEnabledEb(void* qthis, bool arg0);
  // proto:  void QPainter::QPainter(const QPainter & );
extern void* dector_ZN8QPainterC1ERKS_(void* arg0);
extern void _ZN8QPainterC1ERKS_(void* qthis, void* arg0);
  // proto:  void QPainter::drawPoints(const QPolygon & points);
extern void demth_ZN8QPainter10drawPointsERK8QPolygon(void* qthis, void* arg0);
  // proto:  void QPainter::drawPicture(const QPoint & p, const QPicture & picture);
extern void demth_ZN8QPainter11drawPictureERK6QPointRK8QPicture(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::drawRect(int x1, int y1, int w, int h);
extern void demth_ZN8QPainter8drawRectEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3);
  // proto:  void QPainter::drawEllipse(const QRectF & r);
extern void _ZN8QPainter11drawEllipseERK6QRectF(void* qthis, void* arg0);
  // proto:  void QPainter::drawRect(const QRectF & rect);
extern void demth_ZN8QPainter8drawRectERK6QRectF(void* qthis, void* arg0);
  // proto:  void QPainter::drawPoints(const QPointF * points, int pointCount);
extern void _ZN8QPainter10drawPointsEPK7QPointFi(void* qthis, void* arg0, int32_t arg1);
  // proto:  QRegion QPainter::clipRegion();
extern void _ZNK8QPainter10clipRegionEv(void* qthis);
  // proto:  void QPainter::drawText(const QRectF & r, int flags, const QString & text, QRectF * br);
extern void _ZN8QPainter8drawTextERK6QRectFiRK7QStringPS0_(void* qthis, void* arg0, int32_t arg1, void* arg2, void* arg3);
  // proto:  void QPainter::drawLine(const QLineF & line);
extern void demth_ZN8QPainter8drawLineERK6QLineF(void* qthis, void* arg0);
  // proto:  void QPainter::drawLine(const QPointF & p1, const QPointF & p2);
extern void demth_ZN8QPainter8drawLineERK7QPointFS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::drawPixmap(const QRect & r, const QPixmap & pm);
extern void demth_ZN8QPainter10drawPixmapERK5QRectRK7QPixmap(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::drawTiledPixmap(int x, int y, int w, int h, const QPixmap & , int sx, int sy);
extern void demth_ZN8QPainter15drawTiledPixmapEiiiiRK7QPixmapii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, void* arg4, int32_t arg5, int32_t arg6);
  // proto:  QPaintDevice * QPainter::device();
extern void _ZNK8QPainter6deviceEv(void* qthis);
  // proto:  void QPainter::setViewport(const QRect & viewport);
extern void _ZN8QPainter11setViewportERK5QRect(void* qthis, void* arg0);
  // proto:  void QPainter::fillRect(const QRect & , const QColor & color);
extern void _ZN8QPainter8fillRectERK5QRectRK6QColor(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::setBrushOrigin(const QPointF & );
extern void _ZN8QPainter14setBrushOriginERK7QPointF(void* qthis, void* arg0);
  // proto:  void QPainter::drawTextItem(int x, int y, const QTextItem & ti);
extern void demth_ZN8QPainter12drawTextItemEiiRK9QTextItem(void* qthis, int32_t arg0, int32_t arg1, void* arg2);
  // proto:  void QPainter::QPainter(QPaintDevice * );
extern void* dector_ZN8QPainterC1EP12QPaintDevice(void* arg0);
extern void _ZN8QPainterC1EP12QPaintDevice(void* qthis, void* arg0);
  // proto:  void QPainter::drawPixmap(int x, int y, int w, int h, const QPixmap & pm, int sx, int sy, int sw, int sh);
extern void demth_ZN8QPainter10drawPixmapEiiiiRK7QPixmapiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, void* arg4, int32_t arg5, int32_t arg6, int32_t arg7, int32_t arg8);
  // proto:  void QPainter::drawImage(const QPoint & p, const QImage & image);
extern void demth_ZN8QPainter9drawImageERK6QPointRK6QImage(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::drawPie(const QRect & , int a, int alen);
extern void demth_ZN8QPainter7drawPieERK5QRectii(void* qthis, void* arg0, int32_t arg1, int32_t arg2);
  // proto:  void QPainter::drawTextItem(const QPoint & p, const QTextItem & ti);
extern void demth_ZN8QPainter12drawTextItemERK6QPointRK9QTextItem(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::drawLines(const QPoint * pointPairs, int lineCount);
extern void _ZN8QPainter9drawLinesEPK6QPointi(void* qthis, void* arg0, int32_t arg1);
  // proto:  void QPainter::drawPicture(int x, int y, const QPicture & picture);
extern void demth_ZN8QPainter11drawPictureEiiRK8QPicture(void* qthis, int32_t arg0, int32_t arg1, void* arg2);
  // proto:  void QPainter::save();
extern void _ZN8QPainter4saveEv(void* qthis);
  // proto:  void QPainter::translate(qreal dx, qreal dy);
extern void demth_ZN8QPainter9translateEdd(void* qthis, double arg0, double arg1);
  // proto:  QTransform QPainter::combinedTransform();
extern void _ZNK8QPainter17combinedTransformEv(void* qthis);
  // proto:  bool QPainter::end();
extern void _ZN8QPainter3endEv(void* qthis);
  // proto:  void QPainter::setViewport(int x, int y, int w, int h);
extern void demth_ZN8QPainter11setViewportEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3);
  // proto:  void QPainter::drawRoundRect(const QRect & r, int xround, int yround);
extern void demth_ZN8QPainter13drawRoundRectERK5QRectii(void* qthis, void* arg0, int32_t arg1, int32_t arg2);
  // proto:  void QPainter::setWorldTransform(const QTransform & matrix, bool combine);
extern void _ZN8QPainter17setWorldTransformERK10QTransformb(void* qthis, void* arg0, bool arg1);
  // proto:  void QPainter::drawPoints(const QPolygonF & points);
extern void demth_ZN8QPainter10drawPointsERK9QPolygonF(void* qthis, void* arg0);
  // proto:  void QPainter::restore();
extern void _ZN8QPainter7restoreEv(void* qthis);
  // proto:  void QPainter::drawStaticText(const QPoint & topLeftPosition, const QStaticText & staticText);
extern void demth_ZN8QPainter14drawStaticTextERK6QPointRK11QStaticText(void* qthis, void* arg0, void* arg1);
  // proto:  QRectF QPainter::boundingRect(const QRectF & rect, int flags, const QString & text);
extern void _ZN8QPainter12boundingRectERK6QRectFiRK7QString(void* qthis, void* arg0, int32_t arg1, void* arg2);
  // proto:  void QPainter::fillRect(int x, int y, int w, int h, const QBrush & );
extern void demth_ZN8QPainter8fillRectEiiiiRK6QBrush(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, void* arg4);
  // proto:  void QPainter::drawRoundRect(const QRectF & r, int xround, int yround);
extern void _ZN8QPainter13drawRoundRectERK6QRectFii(void* qthis, void* arg0, int32_t arg1, int32_t arg2);
  // proto:  void QPainter::drawPoint(const QPoint & p);
extern void demth_ZN8QPainter9drawPointERK6QPoint(void* qthis, void* arg0);
  // proto: static QPaintDevice * QPainter::redirected(const QPaintDevice * device, QPoint * offset);
extern void _ZN8QPainter10redirectedEPK12QPaintDeviceP6QPoint(void* arg0, void* arg1);
  // proto:  void QPainter::shear(qreal sh, qreal sv);
extern void _ZN8QPainter5shearEdd(void* qthis, double arg0, double arg1);
  // proto:  void QPainter::drawText(const QRect & r, int flags, const QString & text, QRect * br);
extern void _ZN8QPainter8drawTextERK5QRectiRK7QStringPS0_(void* qthis, void* arg0, int32_t arg1, void* arg2, void* arg3);
  // proto:  const QFont & QPainter::font();
extern void _ZNK8QPainter4fontEv(void* qthis);
  // proto:  const QTransform & QPainter::deviceTransform();
extern void _ZNK8QPainter15deviceTransformEv(void* qthis);
  // proto:  void QPainter::eraseRect(int x, int y, int w, int h);
extern void demth_ZN8QPainter9eraseRectEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3);
  // proto:  void QPainter::resetMatrix();
extern void _ZN8QPainter11resetMatrixEv(void* qthis);
  // proto:  void QPainter::drawPolyline(const QPoint * points, int pointCount);
extern void _ZN8QPainter12drawPolylineEPK6QPointi(void* qthis, void* arg0, int32_t arg1);
  // proto:  QPaintEngine * QPainter::paintEngine();
extern void _ZNK8QPainter11paintEngineEv(void* qthis);
  // proto:  void QPainter::drawEllipse(const QRect & r);
extern void _ZN8QPainter11drawEllipseERK5QRect(void* qthis, void* arg0);
  // proto:  void QPainter::drawLine(const QLine & line);
extern void demth_ZN8QPainter8drawLineERK5QLine(void* qthis, void* arg0);
  // proto:  bool QPainter::isActive();
extern void _ZNK8QPainter8isActiveEv(void* qthis);
  // proto:  void QPainter::drawArc(const QRectF & rect, int a, int alen);
extern void _ZN8QPainter7drawArcERK6QRectFii(void* qthis, void* arg0, int32_t arg1, int32_t arg2);
  // proto: static void QPainter::restoreRedirected(const QPaintDevice * device);
extern void _ZN8QPainter17restoreRedirectedEPK12QPaintDevice(void* arg0);
  // proto:  void QPainter::drawPixmap(const QPointF & p, const QPixmap & pm, const QRectF & sr);
extern void demth_ZN8QPainter10drawPixmapERK7QPointFRK7QPixmapRK6QRectF(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QPainter::drawEllipse(const QPointF & center, qreal rx, qreal ry);
extern void demth_ZN8QPainter11drawEllipseERK7QPointFdd(void* qthis, void* arg0, double arg1, double arg2);
  // proto:  void QPainter::drawConvexPolygon(const QPointF * points, int pointCount);
extern void _ZN8QPainter17drawConvexPolygonEPK7QPointFi(void* qthis, void* arg0, int32_t arg1);
  // proto:  void QPainter::setBrushOrigin(const QPoint & );
extern void demth_ZN8QPainter14setBrushOriginERK6QPoint(void* qthis, void* arg0);
  // proto:  void QPainter::drawText(const QRectF & r, const QString & text, const QTextOption & o);
extern void _ZN8QPainter8drawTextERK6QRectFRK7QStringRK11QTextOption(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  bool QPainter::worldMatrixEnabled();
extern void _ZNK8QPainter18worldMatrixEnabledEv(void* qthis);
  // proto:  void QPainter::drawPixmap(const QPoint & p, const QPixmap & pm);
extern void demth_ZN8QPainter10drawPixmapERK6QPointRK7QPixmap(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::drawLine(int x1, int y1, int x2, int y2);
extern void demth_ZN8QPainter8drawLineEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3);
  // proto:  void QPainter::drawPoint(int x, int y);
extern void demth_ZN8QPainter9drawPointEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  const QTransform & QPainter::transform();
extern void _ZNK8QPainter9transformEv(void* qthis);
  // proto: static void QPainter::setRedirected(const QPaintDevice * device, QPaintDevice * replacement, const QPoint & offset);
extern void _ZN8QPainter13setRedirectedEPK12QPaintDevicePS0_RK6QPoint(void* arg0, void* arg1, void* arg2);
  // proto:  void QPainter::drawPixmap(const QRect & targetRect, const QPixmap & pixmap, const QRect & sourceRect);
extern void demth_ZN8QPainter10drawPixmapERK5QRectRK7QPixmapS2_(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  QFontMetrics QPainter::fontMetrics();
extern void _ZNK8QPainter11fontMetricsEv(void* qthis);
  // proto:  void QPainter::drawGlyphRun(const QPointF & position, const QGlyphRun & glyphRun);
extern void _ZN8QPainter12drawGlyphRunERK7QPointFRK9QGlyphRun(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::fillRect(const QRectF & , const QBrush & );
extern void _ZN8QPainter8fillRectERK6QRectFRK6QBrush(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::resetTransform();
extern void _ZN8QPainter14resetTransformEv(void* qthis);
  // proto:  void QPainter::fillRect(const QRect & , const QBrush & );
extern void _ZN8QPainter8fillRectERK5QRectRK6QBrush(void* qthis, void* arg0, void* arg1);
  // proto:  const QBrush & QPainter::brush();
extern void _ZNK8QPainter5brushEv(void* qthis);
  // proto:  void QPainter::~QPainter();
extern void _ZN8QPainterD0Ev(void* qthis);
  // proto:  bool QPainter::begin(QPaintDevice * );
extern void _ZN8QPainter5beginEP12QPaintDevice(void* qthis, void* arg0);
  // proto:  void QPainter::drawRect(const QRect & rect);
extern void demth_ZN8QPainter8drawRectERK5QRect(void* qthis, void* arg0);
  // proto:  void QPainter::drawTextItem(const QPointF & p, const QTextItem & ti);
extern void _ZN8QPainter12drawTextItemERK7QPointFRK9QTextItem(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::scale(qreal sx, qreal sy);
extern void _ZN8QPainter5scaleEdd(void* qthis, double arg0, double arg1);
  // proto:  void QPainter::setWorldMatrix(const QMatrix & matrix, bool combine);
extern void _ZN8QPainter14setWorldMatrixERK7QMatrixb(void* qthis, void* arg0, bool arg1);
  // proto:  QPainterPath QPainter::clipPath();
extern void _ZNK8QPainter8clipPathEv(void* qthis);
  // proto:  QPoint QPainter::brushOrigin();
extern void _ZNK8QPainter11brushOriginEv(void* qthis);
  // proto:  void QPainter::drawConvexPolygon(const QPolygonF & polygon);
extern void demth_ZN8QPainter17drawConvexPolygonERK9QPolygonF(void* qthis, void* arg0);
  // proto:  void QPainter::drawEllipse(int x, int y, int w, int h);
extern void demth_ZN8QPainter11drawEllipseEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3);
  // proto:  void QPainter::drawConvexPolygon(const QPolygon & polygon);
extern void demth_ZN8QPainter17drawConvexPolygonERK8QPolygon(void* qthis, void* arg0);
  // proto:  void QPainter::drawPoints(const QPoint * points, int pointCount);
extern void _ZN8QPainter10drawPointsEPK6QPointi(void* qthis, void* arg0, int32_t arg1);
  // proto:  const QBrush & QPainter::background();
extern void _ZNK8QPainter10backgroundEv(void* qthis);
  // proto:  void QPainter::drawRoundRect(int x, int y, int w, int h, int , int );
extern void demth_ZN8QPainter13drawRoundRectEiiiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4, int32_t arg5);
  // proto:  QRect QPainter::viewport();
extern void _ZNK8QPainter8viewportEv(void* qthis);
  // proto:  void QPainter::drawArc(const QRect & , int a, int alen);
extern void demth_ZN8QPainter7drawArcERK5QRectii(void* qthis, void* arg0, int32_t arg1, int32_t arg2);
  // proto:  void QPainter::fillPath(const QPainterPath & path, const QBrush & brush);
extern void _ZN8QPainter8fillPathERK12QPainterPathRK6QBrush(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::drawText(int x, int y, int w, int h, int flags, const QString & text, QRect * br);
extern void demth_ZN8QPainter8drawTextEiiiiiRK7QStringP5QRect(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4, void* arg5, void* arg6);
  // proto:  bool QPainter::matrixEnabled();
extern void _ZNK8QPainter13matrixEnabledEv(void* qthis);
  // proto:  void QPainter::drawPolyline(const QPointF * points, int pointCount);
extern void _ZN8QPainter12drawPolylineEPK7QPointFi(void* qthis, void* arg0, int32_t arg1);
  // proto:  void QPainter::setTransform(const QTransform & transform, bool combine);
extern void _ZN8QPainter12setTransformERK10QTransformb(void* qthis, void* arg0, bool arg1);
  // proto:  void QPainter::setPen(const QColor & color);
extern void _ZN8QPainter6setPenERK6QColor(void* qthis, void* arg0);
  // proto:  void QPainter::eraseRect(const QRect & );
extern void demth_ZN8QPainter9eraseRectERK5QRect(void* qthis, void* arg0);
  // proto:  QRect QPainter::window();
extern void _ZNK8QPainter6windowEv(void* qthis);
  // proto:  void QPainter::drawImage(const QRect & r, const QImage & image);
extern void demth_ZN8QPainter9drawImageERK5QRectRK6QImage(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::initFrom(const QPaintDevice * device);
extern void _ZN8QPainter8initFromEPK12QPaintDevice(void* qthis, void* arg0);
  // proto:  QFontInfo QPainter::fontInfo();
extern void _ZNK8QPainter8fontInfoEv(void* qthis);
  // proto:  void QPainter::endNativePainting();
extern void _ZN8QPainter17endNativePaintingEv(void* qthis);
  // proto:  void QPainter::setViewTransformEnabled(bool enable);
extern void _ZN8QPainter23setViewTransformEnabledEb(void* qthis, bool arg0);
  // proto:  void QPainter::drawPoint(const QPointF & pt);
extern void demth_ZN8QPainter9drawPointERK7QPointF(void* qthis, void* arg0);
  // proto:  void QPainter::setOpacity(qreal opacity);
extern void _ZN8QPainter10setOpacityEd(void* qthis, double arg0);
  // proto:  void QPainter::fillRect(const QRectF & , const QColor & color);
extern void _ZN8QPainter8fillRectERK6QRectFRK6QColor(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::QPainter();
extern void* dector_ZN8QPainterC1Ev();
extern void _ZN8QPainterC1Ev(void* qthis);
  // proto:  void QPainter::translate(const QPointF & offset);
extern void _ZN8QPainter9translateERK7QPointF(void* qthis, void* arg0);
  // proto:  void QPainter::drawText(const QPointF & p, const QString & s);
extern void _ZN8QPainter8drawTextERK7QPointFRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::drawImage(const QPointF & p, const QImage & image);
extern void _ZN8QPainter9drawImageERK7QPointFRK6QImage(void* qthis, void* arg0, void* arg1);
  // proto:  const QPen & QPainter::pen();
extern void _ZNK8QPainter3penEv(void* qthis);
  // proto:  void QPainter::rotate(qreal a);
extern void _ZN8QPainter6rotateEd(void* qthis, double arg0);
  // proto:  QRectF QPainter::clipBoundingRect();
extern void _ZNK8QPainter16clipBoundingRectEv(void* qthis);
  // proto:  void QPainter::drawLine(const QPoint & p1, const QPoint & p2);
extern void demth_ZN8QPainter8drawLineERK6QPointS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::drawPie(const QRectF & rect, int a, int alen);
extern void _ZN8QPainter7drawPieERK6QRectFii(void* qthis, void* arg0, int32_t arg1, int32_t arg2);
  // proto:  void QPainter::drawText(const QPoint & p, const QString & s);
extern void demth_ZN8QPainter8drawTextERK6QPointRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  void QPainter::setWindow(int x, int y, int w, int h);
extern void demth_ZN8QPainter9setWindowEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3);
  // proto:  void QPainter::beginNativePainting();
extern void _ZN8QPainter19beginNativePaintingEv(void* qthis);
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

// class sizeof(QPainter)=1
type QPainter struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  QRectF QPainter::boundingRect(const QRectF & rect, const QString & text, const QTextOption & o);
func (this *QPainter) boundingRect(args ...interface{}) () {
  // boundingRect(const class QRectF &, const class QString &, const class QTextOption &)
  // boundingRect(const class QRect &, int, const class QString &)
  // boundingRect(int, int, int, int, int, const class QString &)
  // boundingRect(const class QRectF &, int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QTextOption{}) // "const QTextOption &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = qtrt.Int32Ty(false) // "int"
  vtys[2][4] = qtrt.Int32Ty(false) // "int"
  vtys[2][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter12boundingRectERK6QRectFRK7QStringRK11QTextOption
    // invoke: QRectF boundingRect(const class QRectF &, const class QString &, const class QTextOption &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QTextOption).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN8QPainter12boundingRectERK6QRectFRK7QStringRK11QTextOption(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QPainter12boundingRectERK5QRectiRK7QString
    // invoke: QRect boundingRect(const class QRect &, int, const class QString &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN8QPainter12boundingRectERK5QRectiRK7QString(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN8QPainter12boundingRectEiiiiiRK7QString
    // invoke: QRect boundingRect(int, int, int, int, int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(QString).qclsinst
    if false {fmt.Println(arg5)}
    C.demth_ZN8QPainter12boundingRectEiiiiiRK7QString(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  case 3:
    // invoke: _ZN8QPainter12boundingRectERK6QRectFiRK7QString
    // invoke: QRectF boundingRect(const class QRectF &, int, const class QString &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN8QPainter12boundingRectERK6QRectFiRK7QString(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPainter", "boundingRect", args)
  }

}

  // proto:  void QPainter::drawPicture(const QPointF & p, const QPicture & picture);
func (this *QPainter) drawPicture(args ...interface{}) () {
  // drawPicture(const class QPointF &, const class QPicture &)
  // drawPicture(const class QPoint &, const class QPicture &)
  // drawPicture(int, int, const class QPicture &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QPicture{}) // "const QPicture &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1][1] = reflect.TypeOf(QPicture{}) // "const QPicture &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = reflect.TypeOf(QPicture{}) // "const QPicture &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter11drawPictureERK7QPointFRK8QPicture
    // invoke: void drawPicture(const class QPointF &, const class QPicture &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPicture).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QPainter11drawPictureERK7QPointFRK8QPicture(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPainter11drawPictureERK6QPointRK8QPicture
    // invoke: void drawPicture(const class QPoint &, const class QPicture &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPicture).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN8QPainter11drawPictureERK6QPointRK8QPicture(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN8QPainter11drawPictureEiiRK8QPicture
    // invoke: void drawPicture(int, int, const class QPicture &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPicture).qclsinst
    if false {fmt.Println(arg2)}
    C.demth_ZN8QPainter11drawPictureEiiRK8QPicture(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPainter", "drawPicture", args)
  }

}

  // proto:  const QMatrix & QPainter::worldMatrix();
func (this *QPainter) worldMatrix(args ...interface{}) () {
  // worldMatrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter11worldMatrixEv
    // invoke: const QMatrix & worldMatrix()
    C._ZNK8QPainter11worldMatrixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "worldMatrix", args)
  }

}

  // proto:  void QPainter::drawText(const QPointF & p, const QString & str, int tf, int justificationPadding);
func (this *QPainter) drawText(args ...interface{}) () {
  // drawText(const class QPointF &, const class QString &, int, int)
  // drawText(int, int, const class QString &)
  // drawText(const class QRectF &, int, const class QString &, class QRectF *)
  // drawText(const class QRect &, int, const class QString &, class QRect *)
  // drawText(const class QRectF &, const class QString &, const class QTextOption &)
  // drawText(int, int, int, int, int, const class QString &, class QRect *)
  // drawText(const class QPointF &, const class QString &)
  // drawText(const class QPoint &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][3] = reflect.TypeOf(QRectF{}) // "QRectF *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][3] = reflect.TypeOf(QRect{}) // "QRect *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[4][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4][2] = reflect.TypeOf(QTextOption{}) // "const QTextOption &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[5][2] = qtrt.Int32Ty(false) // "int"
  vtys[5][3] = qtrt.Int32Ty(false) // "int"
  vtys[5][4] = qtrt.Int32Ty(false) // "int"
  vtys[5][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[5][6] = reflect.TypeOf(QRect{}) // "QRect *"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[6][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[7][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8drawTextERK7QPointFRK7QStringii
    // invoke: void drawText(const class QPointF &, const class QString &, int, int)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN8QPainter8drawTextERK7QPointFRK7QStringii(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN8QPainter8drawTextEiiRK7QString
    // invoke: void drawText(int, int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    C.demth_ZN8QPainter8drawTextEiiRK7QString(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN8QPainter8drawTextERK6QRectFiRK7QStringPS0_
    // invoke: void drawText(const class QRectF &, int, const class QString &, class QRectF *)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QRectF).qclsinst
    if false {fmt.Println(arg3)}
    C._ZN8QPainter8drawTextERK6QRectFiRK7QStringPS0_(this.qclsinst, arg0, arg1, arg2, arg3)
  case 3:
    // invoke: _ZN8QPainter8drawTextERK5QRectiRK7QStringPS0_
    // invoke: void drawText(const class QRect &, int, const class QString &, class QRect *)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QRect).qclsinst
    if false {fmt.Println(arg3)}
    C._ZN8QPainter8drawTextERK5QRectiRK7QStringPS0_(this.qclsinst, arg0, arg1, arg2, arg3)
  case 4:
    // invoke: _ZN8QPainter8drawTextERK6QRectFRK7QStringRK11QTextOption
    // invoke: void drawText(const class QRectF &, const class QString &, const class QTextOption &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QTextOption).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN8QPainter8drawTextERK6QRectFRK7QStringRK11QTextOption(this.qclsinst, arg0, arg1, arg2)
  case 5:
    // invoke: _ZN8QPainter8drawTextEiiiiiRK7QStringP5QRect
    // invoke: void drawText(int, int, int, int, int, const class QString &, class QRect *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(QString).qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = args[6].(QRect).qclsinst
    if false {fmt.Println(arg6)}
    C.demth_ZN8QPainter8drawTextEiiiiiRK7QStringP5QRect(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
  case 6:
    // invoke: _ZN8QPainter8drawTextERK7QPointFRK7QString
    // invoke: void drawText(const class QPointF &, const class QString &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QPainter8drawTextERK7QPointFRK7QString(this.qclsinst, arg0, arg1)
  case 7:
    // invoke: _ZN8QPainter8drawTextERK6QPointRK7QString
    // invoke: void drawText(const class QPoint &, const class QString &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN8QPainter8drawTextERK6QPointRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawText", args)
  }

}

  // proto:  void QPainter::fillRect(int x, int y, int w, int h, const QColor & color);
func (this *QPainter) fillRect(args ...interface{}) () {
  // fillRect(const class QRect &, Qt::GlobalColor)
  // fillRect(int, int, int, int, const class QColor &)
  // fillRect(int, int, int, int, Qt::BrushStyle)
  // fillRect(const class QRectF &, Qt::BrushStyle)
  // fillRect(const class QRectF &, Qt::GlobalColor)
  // fillRect(int, int, int, int, Qt::GlobalColor)
  // fillRect(const class QRect &, const class QColor &)
  // fillRect(int, int, int, int, const class QBrush &)
  // fillRect(const class QRectF &, const class QBrush &)
  // fillRect(const class QRect &, const class QBrush &)
  // fillRect(const class QRect &, Qt::BrushStyle)
  // fillRect(const class QRectF &, const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[0][1] = qtrt.Int32Ty(false) // "Qt::GlobalColor"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = qtrt.Int32Ty(false) // "int"
  vtys[2][4] = qtrt.Int32Ty(false) // "Qt::BrushStyle"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[3][1] = qtrt.Int32Ty(false) // "Qt::BrushStyle"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[4][1] = qtrt.Int32Ty(false) // "Qt::GlobalColor"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[5][2] = qtrt.Int32Ty(false) // "int"
  vtys[5][3] = qtrt.Int32Ty(false) // "int"
  vtys[5][4] = qtrt.Int32Ty(false) // "Qt::GlobalColor"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[6][1] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"
  vtys[7][2] = qtrt.Int32Ty(false) // "int"
  vtys[7][3] = qtrt.Int32Ty(false) // "int"
  vtys[7][4] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[8][1] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[9][1] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[10][1] = qtrt.Int32Ty(false) // "Qt::BrushStyle"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[11][1] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8fillRectEiiiiRK6QColor
    // invoke: void fillRect(int, int, int, int, const class QColor &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QColor).qclsinst
    if false {fmt.Println(arg4)}
    C.demth_ZN8QPainter8fillRectEiiiiRK6QColor(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 1:
    // invoke: _ZN8QPainter8fillRectERK5QRectRK6QColor
    // invoke: void fillRect(const class QRect &, const class QColor &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QColor).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QPainter8fillRectERK5QRectRK6QColor(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN8QPainter8fillRectEiiiiRK6QBrush
    // invoke: void fillRect(int, int, int, int, const class QBrush &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QBrush).qclsinst
    if false {fmt.Println(arg4)}
    C.demth_ZN8QPainter8fillRectEiiiiRK6QBrush(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 3:
    // invoke: _ZN8QPainter8fillRectERK6QRectFRK6QBrush
    // invoke: void fillRect(const class QRectF &, const class QBrush &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QBrush).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QPainter8fillRectERK6QRectFRK6QBrush(this.qclsinst, arg0, arg1)
  case 4:
    // invoke: _ZN8QPainter8fillRectERK5QRectRK6QBrush
    // invoke: void fillRect(const class QRect &, const class QBrush &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QBrush).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QPainter8fillRectERK5QRectRK6QBrush(this.qclsinst, arg0, arg1)
  case 5:
    // invoke: _ZN8QPainter8fillRectERK6QRectFRK6QColor
    // invoke: void fillRect(const class QRectF &, const class QColor &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QColor).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QPainter8fillRectERK6QRectFRK6QColor(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "fillRect", args)
  }

}

  // proto:  const QMatrix & QPainter::matrix();
func (this *QPainter) matrix(args ...interface{}) () {
  // matrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter6matrixEv
    // invoke: const QMatrix & matrix()
    C._ZNK8QPainter6matrixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "matrix", args)
  }

}

  // proto:  qreal QPainter::opacity();
func (this *QPainter) opacity(args ...interface{}) () {
  // opacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter7opacityEv
    // invoke: qreal opacity()
    C._ZNK8QPainter7opacityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "opacity", args)
  }

}

  // proto:  void QPainter::drawTiledPixmap(const QRectF & rect, const QPixmap & pm, const QPointF & offset);
func (this *QPainter) drawTiledPixmap(args ...interface{}) () {
  // drawTiledPixmap(const class QRectF &, const class QPixmap &, const class QPointF &)
  // drawTiledPixmap(const class QRect &, const class QPixmap &, const class QPoint &)
  // drawTiledPixmap(int, int, int, int, const class QPixmap &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[0][2] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[1][2] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = qtrt.Int32Ty(false) // "int"
  vtys[2][4] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[2][5] = qtrt.Int32Ty(false) // "int"
  vtys[2][6] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF
    // invoke: void drawTiledPixmap(const class QRectF &, const class QPixmap &, const class QPointF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPixmap).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPointF).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN8QPainter15drawTiledPixmapERK6QRectFRK7QPixmapRK7QPointF(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QPainter15drawTiledPixmapERK5QRectRK7QPixmapRK6QPoint
    // invoke: void drawTiledPixmap(const class QRect &, const class QPixmap &, const class QPoint &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPixmap).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPoint).qclsinst
    if false {fmt.Println(arg2)}
    C.demth_ZN8QPainter15drawTiledPixmapERK5QRectRK7QPixmapRK6QPoint(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN8QPainter15drawTiledPixmapEiiiiRK7QPixmapii
    // invoke: void drawTiledPixmap(int, int, int, int, const class QPixmap &, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QPixmap).qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(args[6].(int32))
    if false {fmt.Println(arg6)}
    C.demth_ZN8QPainter15drawTiledPixmapEiiiiRK7QPixmapii(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
  default:
    qtrt.ErrorResolve("QPainter", "drawTiledPixmap", args)
  }

}

  // proto:  void QPainter::setBackground(const QBrush & bg);
func (this *QPainter) setBackground(args ...interface{}) () {
  // setBackground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter13setBackgroundERK6QBrush
    // invoke: void setBackground(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QPainter13setBackgroundERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setBackground", args)
  }

}

  // proto:  void QPainter::drawChord(const QRectF & rect, int a, int alen);
func (this *QPainter) drawChord(args ...interface{}) () {
  // drawChord(const class QRectF &, int, int)
  // drawChord(const class QRect &, int, int)
  // drawChord(int, int, int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = qtrt.Int32Ty(false) // "int"
  vtys[2][4] = qtrt.Int32Ty(false) // "int"
  vtys[2][5] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9drawChordERK6QRectFii
    // invoke: void drawChord(const class QRectF &, int, int)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN8QPainter9drawChordERK6QRectFii(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QPainter9drawChordERK5QRectii
    // invoke: void drawChord(const class QRect &, int, int)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.demth_ZN8QPainter9drawChordERK5QRectii(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN8QPainter9drawChordEiiiiii
    // invoke: void drawChord(int, int, int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    C.demth_ZN8QPainter9drawChordEiiiiii(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  default:
    qtrt.ErrorResolve("QPainter", "drawChord", args)
  }

}

  // proto:  void QPainter::drawImage(const QRectF & r, const QImage & image);
func (this *QPainter) drawImage(args ...interface{}) () {
  // drawImage(const class QRectF &, const class QImage &)
  // drawImage(const class QPointF &, const class QImage &, const class QRectF &, Qt::ImageConversionFlags)
  // drawImage(const class QRect &, const class QImage &, const class QRect &, Qt::ImageConversionFlags)
  // drawImage(int, int, const class QImage &, int, int, int, int, Qt::ImageConversionFlags)
  // drawImage(const class QPoint &, const class QImage &, const class QRect &, Qt::ImageConversionFlags)
  // drawImage(const class QPoint &, const class QImage &)
  // drawImage(const class QRect &, const class QImage &)
  // drawImage(const class QPointF &, const class QImage &)
  // drawImage(const class QRectF &, const class QImage &, const class QRectF &, Qt::ImageConversionFlags)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[1][2] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][3] = qtrt.Int64Ty(false) // "Qt::ImageConversionFlags"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[2][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[2][2] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[2][3] = qtrt.Int64Ty(false) // "Qt::ImageConversionFlags"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[3][3] = qtrt.Int32Ty(false) // "int"
  vtys[3][4] = qtrt.Int32Ty(false) // "int"
  vtys[3][5] = qtrt.Int32Ty(false) // "int"
  vtys[3][6] = qtrt.Int32Ty(false) // "int"
  vtys[3][7] = qtrt.Int64Ty(false) // "Qt::ImageConversionFlags"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[4][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[4][2] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[4][3] = qtrt.Int64Ty(false) // "Qt::ImageConversionFlags"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[5][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[6][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[7][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[8][1] = reflect.TypeOf(QImage{}) // "const QImage &"
  vtys[8][2] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[8][3] = qtrt.Int64Ty(false) // "Qt::ImageConversionFlags"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9drawImageERK6QRectFRK6QImage
    // invoke: void drawImage(const class QRectF &, const class QImage &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QImage).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN8QPainter9drawImageERK6QRectFRK6QImage(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPainter9drawImageERK6QPointRK6QImage
    // invoke: void drawImage(const class QPoint &, const class QImage &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QImage).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN8QPainter9drawImageERK6QPointRK6QImage(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN8QPainter9drawImageERK5QRectRK6QImage
    // invoke: void drawImage(const class QRect &, const class QImage &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QImage).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN8QPainter9drawImageERK5QRectRK6QImage(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN8QPainter9drawImageERK7QPointFRK6QImage
    // invoke: void drawImage(const class QPointF &, const class QImage &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QImage).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QPainter9drawImageERK7QPointFRK6QImage(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawImage", args)
  }

}

  // proto:  void QPainter::setClipping(bool enable);
func (this *QPainter) setClipping(args ...interface{}) () {
  // setClipping(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter11setClippingEb
    // invoke: void setClipping(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN8QPainter11setClippingEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setClipping", args)
  }

}

  // proto:  void QPainter::setBrush(const QBrush & brush);
func (this *QPainter) setBrush(args ...interface{}) () {
  // setBrush(const class QBrush &)
  // setBrush(Qt::BrushStyle)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "Qt::BrushStyle"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8setBrushERK6QBrush
    // invoke: void setBrush(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QPainter8setBrushERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setBrush", args)
  }

}

  // proto:  void QPainter::setMatrix(const QMatrix & matrix, bool combine);
func (this *QPainter) setMatrix(args ...interface{}) () {
  // setMatrix(const class QMatrix &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9setMatrixERK7QMatrixb
    // invoke: void setMatrix(const class QMatrix &, _Bool)
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZN8QPainter9setMatrixERK7QMatrixb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "setMatrix", args)
  }

}

  // proto:  void QPainter::eraseRect(const QRectF & );
func (this *QPainter) eraseRect(args ...interface{}) () {
  // eraseRect(const class QRectF &)
  // eraseRect(int, int, int, int)
  // eraseRect(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9eraseRectERK6QRectF
    // invoke: void eraseRect(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QPainter9eraseRectERK6QRectF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN8QPainter9eraseRectEiiii
    // invoke: void eraseRect(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.demth_ZN8QPainter9eraseRectEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  case 2:
    // invoke: _ZN8QPainter9eraseRectERK5QRect
    // invoke: void eraseRect(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPainter9eraseRectERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "eraseRect", args)
  }

}

  // proto:  void QPainter::translate(const QPoint & offset);
func (this *QPainter) translate(args ...interface{}) () {
  // translate(const class QPoint &)
  // translate(qreal, qreal)
  // translate(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9translateERK6QPoint
    // invoke: void translate(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPainter9translateERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN8QPainter9translateEdd
    // invoke: void translate(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.demth_ZN8QPainter9translateEdd(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN8QPainter9translateERK7QPointF
    // invoke: void translate(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QPainter9translateERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "translate", args)
  }

}

  // proto:  bool QPainter::viewTransformEnabled();
func (this *QPainter) viewTransformEnabled(args ...interface{}) () {
  // viewTransformEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter20viewTransformEnabledEv
    // invoke: bool viewTransformEnabled()
    C._ZNK8QPainter20viewTransformEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "viewTransformEnabled", args)
  }

}

  // proto:  void QPainter::setPen(const QPen & pen);
func (this *QPainter) setPen(args ...interface{}) () {
  // setPen(const class QPen &)
  // setPen(Qt::PenStyle)
  // setPen(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPen{}) // "const QPen &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "Qt::PenStyle"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter6setPenERK4QPen
    // invoke: void setPen(const class QPen &)
    var arg0 = args[0].(QPen).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QPainter6setPenERK4QPen(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN8QPainter6setPenERK6QColor
    // invoke: void setPen(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QPainter6setPenERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setPen", args)
  }

}

  // proto:  void QPainter::drawLines(const QLineF * lines, int lineCount);
func (this *QPainter) drawLines(args ...interface{}) () {
  // drawLines(const class QLineF *, int)
  // drawLines(const class QPointF *, int)
  // drawLines(const class QLine *, int)
  // drawLines(const QVector<class QLine> &)
  // drawLines(const class QPoint *, int)
  // drawLines(const QVector<class QLineF> &)
  // drawLines(const QVector<class QPoint> &)
  // drawLines(const QVector<class QPointF> &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLineF{}) // "const QLineF *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QLine{}) // "const QLine *"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  // vtys[3][0] = reflect.TypeOf(QVector<QLine>{}) // "const QVector<QLine> &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPoint{}) // "const QPoint *"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  // vtys[5][0] = reflect.TypeOf(QVector<QLineF>{}) // "const QVector<QLineF> &"
  vtys[6] = make(map[int32]reflect.Type)
  // vtys[6][0] = reflect.TypeOf(QVector<QPoint>{}) // "const QVector<QPoint> &"
  vtys[7] = make(map[int32]reflect.Type)
  // vtys[7][0] = reflect.TypeOf(QVector<QPointF>{}) // "const QVector<QPointF> &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9drawLinesEPK6QLineFi
    // invoke: void drawLines(const class QLineF *, int)
    var arg0 = args[0].(QLineF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN8QPainter9drawLinesEPK6QLineFi(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPainter9drawLinesEPK7QPointFi
    // invoke: void drawLines(const class QPointF *, int)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN8QPainter9drawLinesEPK7QPointFi(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN8QPainter9drawLinesEPK5QLinei
    // invoke: void drawLines(const class QLine *, int)
    var arg0 = args[0].(QLine).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN8QPainter9drawLinesEPK5QLinei(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN8QPainter9drawLinesEPK6QPointi
    // invoke: void drawLines(const class QPoint *, int)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN8QPainter9drawLinesEPK6QPointi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawLines", args)
  }

}

  // proto:  void QPainter::setBrushOrigin(int x, int y);
func (this *QPainter) setBrushOrigin(args ...interface{}) () {
  // setBrushOrigin(int, int)
  // setBrushOrigin(const class QPointF &)
  // setBrushOrigin(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter14setBrushOriginEii
    // invoke: void setBrushOrigin(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZN8QPainter14setBrushOriginEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPainter14setBrushOriginERK7QPointF
    // invoke: void setBrushOrigin(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QPainter14setBrushOriginERK7QPointF(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN8QPainter14setBrushOriginERK6QPoint
    // invoke: void setBrushOrigin(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPainter14setBrushOriginERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setBrushOrigin", args)
  }

}

  // proto:  const QTransform & QPainter::worldTransform();
func (this *QPainter) worldTransform(args ...interface{}) () {
  // worldTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter14worldTransformEv
    // invoke: const QTransform & worldTransform()
    C._ZNK8QPainter14worldTransformEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "worldTransform", args)
  }

}

  // proto:  void QPainter::drawRects(const QRect * rects, int rectCount);
func (this *QPainter) drawRects(args ...interface{}) () {
  // drawRects(const QVector<class QRect> &)
  // drawRects(const class QRect *, int)
  // drawRects(const class QRectF *, int)
  // drawRects(const QVector<class QRectF> &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  // vtys[0][0] = reflect.TypeOf(QVector<QRect>{}) // "const QVector<QRect> &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRectF{}) // "const QRectF *"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  // vtys[3][0] = reflect.TypeOf(QVector<QRectF>{}) // "const QVector<QRectF> &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9drawRectsEPK5QRecti
    // invoke: void drawRects(const class QRect *, int)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN8QPainter9drawRectsEPK5QRecti(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPainter9drawRectsEPK6QRectFi
    // invoke: void drawRects(const class QRectF *, int)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN8QPainter9drawRectsEPK6QRectFi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawRects", args)
  }

}

  // proto:  void QPainter::drawEllipse(const QPoint & center, int rx, int ry);
func (this *QPainter) drawEllipse(args ...interface{}) () {
  // drawEllipse(const class QPoint &, int, int)
  // drawEllipse(const class QRectF &)
  // drawEllipse(const class QRect &)
  // drawEllipse(const class QPointF &, qreal, qreal)
  // drawEllipse(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[3][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[3][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[4][2] = qtrt.Int32Ty(false) // "int"
  vtys[4][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter11drawEllipseERK6QPointii
    // invoke: void drawEllipse(const class QPoint &, int, int)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.demth_ZN8QPainter11drawEllipseERK6QPointii(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QPainter11drawEllipseERK6QRectF
    // invoke: void drawEllipse(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QPainter11drawEllipseERK6QRectF(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN8QPainter11drawEllipseERK5QRect
    // invoke: void drawEllipse(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QPainter11drawEllipseERK5QRect(this.qclsinst, arg0)
  case 3:
    // invoke: _ZN8QPainter11drawEllipseERK7QPointFdd
    // invoke: void drawEllipse(const class QPointF &, qreal, qreal)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    C.demth_ZN8QPainter11drawEllipseERK7QPointFdd(this.qclsinst, arg0, arg1, arg2)
  case 4:
    // invoke: _ZN8QPainter11drawEllipseEiiii
    // invoke: void drawEllipse(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.demth_ZN8QPainter11drawEllipseEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPainter", "drawEllipse", args)
  }

}

  // proto:  void QPainter::drawArc(int x, int y, int w, int h, int a, int alen);
func (this *QPainter) drawArc(args ...interface{}) () {
  // drawArc(int, int, int, int, int, int)
  // drawArc(const class QRectF &, int, int)
  // drawArc(const class QRect &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[0][5] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter7drawArcEiiiiii
    // invoke: void drawArc(int, int, int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    C.demth_ZN8QPainter7drawArcEiiiiii(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  case 1:
    // invoke: _ZN8QPainter7drawArcERK6QRectFii
    // invoke: void drawArc(const class QRectF &, int, int)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN8QPainter7drawArcERK6QRectFii(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN8QPainter7drawArcERK5QRectii
    // invoke: void drawArc(const class QRect &, int, int)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.demth_ZN8QPainter7drawArcERK5QRectii(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPainter", "drawArc", args)
  }

}

  // proto:  void QPainter::drawPolyline(const QPolygonF & polyline);
func (this *QPainter) drawPolyline(args ...interface{}) () {
  // drawPolyline(const class QPolygonF &)
  // drawPolyline(const class QPolygon &)
  // drawPolyline(const class QPoint *, int)
  // drawPolyline(const class QPointF *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPoint{}) // "const QPoint *"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPointF{}) // "const QPointF *"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter12drawPolylineERK9QPolygonF
    // invoke: void drawPolyline(const class QPolygonF &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPainter12drawPolylineERK9QPolygonF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN8QPainter12drawPolylineERK8QPolygon
    // invoke: void drawPolyline(const class QPolygon &)
    var arg0 = args[0].(QPolygon).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPainter12drawPolylineERK8QPolygon(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN8QPainter12drawPolylineEPK6QPointi
    // invoke: void drawPolyline(const class QPoint *, int)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN8QPainter12drawPolylineEPK6QPointi(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN8QPainter12drawPolylineEPK7QPointFi
    // invoke: void drawPolyline(const class QPointF *, int)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN8QPainter12drawPolylineEPK7QPointFi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawPolyline", args)
  }

}

  // proto:  bool QPainter::hasClipping();
func (this *QPainter) hasClipping(args ...interface{}) () {
  // hasClipping()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter11hasClippingEv
    // invoke: bool hasClipping()
    C._ZNK8QPainter11hasClippingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "hasClipping", args)
  }

}

  // proto:  void QPainter::drawPixmap(const QRectF & targetRect, const QPixmap & pixmap, const QRectF & sourceRect);
func (this *QPainter) drawPixmap(args ...interface{}) () {
  // drawPixmap(const class QRectF &, const class QPixmap &, const class QRectF &)
  // drawPixmap(int, int, const class QPixmap &, int, int, int, int)
  // drawPixmap(int, int, const class QPixmap &)
  // drawPixmap(int, int, int, int, const class QPixmap &)
  // drawPixmap(const class QPointF &, const class QPixmap &)
  // drawPixmap(const class QPoint &, const class QPixmap &, const class QRect &)
  // drawPixmap(const class QRect &, const class QPixmap &)
  // drawPixmap(int, int, int, int, const class QPixmap &, int, int, int, int)
  // drawPixmap(const class QPointF &, const class QPixmap &, const class QRectF &)
  // drawPixmap(const class QPoint &, const class QPixmap &)
  // drawPixmap(const class QRect &, const class QPixmap &, const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[0][2] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.Int32Ty(false) // "int"
  vtys[1][5] = qtrt.Int32Ty(false) // "int"
  vtys[1][6] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[3][3] = qtrt.Int32Ty(false) // "int"
  vtys[3][4] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[4][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[5][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[5][2] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[6][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"
  vtys[7][2] = qtrt.Int32Ty(false) // "int"
  vtys[7][3] = qtrt.Int32Ty(false) // "int"
  vtys[7][4] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[7][5] = qtrt.Int32Ty(false) // "int"
  vtys[7][6] = qtrt.Int32Ty(false) // "int"
  vtys[7][7] = qtrt.Int32Ty(false) // "int"
  vtys[7][8] = qtrt.Int32Ty(false) // "int"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[8][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[8][2] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[9][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[10][1] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[10][2] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter10drawPixmapERK6QRectFRK7QPixmapS2_
    // invoke: void drawPixmap(const class QRectF &, const class QPixmap &, const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPixmap).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QRectF).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN8QPainter10drawPixmapERK6QRectFRK7QPixmapS2_(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QPainter10drawPixmapEiiRK7QPixmapiiii
    // invoke: void drawPixmap(int, int, const class QPixmap &, int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPixmap).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(args[6].(int32))
    if false {fmt.Println(arg6)}
    C.demth_ZN8QPainter10drawPixmapEiiRK7QPixmapiiii(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
  case 2:
    // invoke: _ZN8QPainter10drawPixmapEiiRK7QPixmap
    // invoke: void drawPixmap(int, int, const class QPixmap &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPixmap).qclsinst
    if false {fmt.Println(arg2)}
    C.demth_ZN8QPainter10drawPixmapEiiRK7QPixmap(this.qclsinst, arg0, arg1, arg2)
  case 3:
    // invoke: _ZN8QPainter10drawPixmapEiiiiRK7QPixmap
    // invoke: void drawPixmap(int, int, int, int, const class QPixmap &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QPixmap).qclsinst
    if false {fmt.Println(arg4)}
    C.demth_ZN8QPainter10drawPixmapEiiiiRK7QPixmap(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 4:
    // invoke: _ZN8QPainter10drawPixmapERK7QPointFRK7QPixmap
    // invoke: void drawPixmap(const class QPointF &, const class QPixmap &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPixmap).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QPainter10drawPixmapERK7QPointFRK7QPixmap(this.qclsinst, arg0, arg1)
  case 5:
    // invoke: _ZN8QPainter10drawPixmapERK6QPointRK7QPixmapRK5QRect
    // invoke: void drawPixmap(const class QPoint &, const class QPixmap &, const class QRect &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPixmap).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QRect).qclsinst
    if false {fmt.Println(arg2)}
    C.demth_ZN8QPainter10drawPixmapERK6QPointRK7QPixmapRK5QRect(this.qclsinst, arg0, arg1, arg2)
  case 6:
    // invoke: _ZN8QPainter10drawPixmapERK5QRectRK7QPixmap
    // invoke: void drawPixmap(const class QRect &, const class QPixmap &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPixmap).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN8QPainter10drawPixmapERK5QRectRK7QPixmap(this.qclsinst, arg0, arg1)
  case 7:
    // invoke: _ZN8QPainter10drawPixmapEiiiiRK7QPixmapiiii
    // invoke: void drawPixmap(int, int, int, int, const class QPixmap &, int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QPixmap).qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(args[6].(int32))
    if false {fmt.Println(arg6)}
    var arg7 = C.int32_t(args[7].(int32))
    if false {fmt.Println(arg7)}
    var arg8 = C.int32_t(args[8].(int32))
    if false {fmt.Println(arg8)}
    C.demth_ZN8QPainter10drawPixmapEiiiiRK7QPixmapiiii(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
  case 8:
    // invoke: _ZN8QPainter10drawPixmapERK7QPointFRK7QPixmapRK6QRectF
    // invoke: void drawPixmap(const class QPointF &, const class QPixmap &, const class QRectF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPixmap).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QRectF).qclsinst
    if false {fmt.Println(arg2)}
    C.demth_ZN8QPainter10drawPixmapERK7QPointFRK7QPixmapRK6QRectF(this.qclsinst, arg0, arg1, arg2)
  case 9:
    // invoke: _ZN8QPainter10drawPixmapERK6QPointRK7QPixmap
    // invoke: void drawPixmap(const class QPoint &, const class QPixmap &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPixmap).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN8QPainter10drawPixmapERK6QPointRK7QPixmap(this.qclsinst, arg0, arg1)
  case 10:
    // invoke: _ZN8QPainter10drawPixmapERK5QRectRK7QPixmapS2_
    // invoke: void drawPixmap(const class QRect &, const class QPixmap &, const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPixmap).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QRect).qclsinst
    if false {fmt.Println(arg2)}
    C.demth_ZN8QPainter10drawPixmapERK5QRectRK7QPixmapS2_(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPainter", "drawPixmap", args)
  }

}

  // proto:  void QPainter::drawStaticText(int left, int top, const QStaticText & staticText);
func (this *QPainter) drawStaticText(args ...interface{}) () {
  // drawStaticText(int, int, const class QStaticText &)
  // drawStaticText(const class QPointF &, const class QStaticText &)
  // drawStaticText(const class QPoint &, const class QStaticText &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QStaticText{}) // "const QStaticText &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1][1] = reflect.TypeOf(QStaticText{}) // "const QStaticText &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[2][1] = reflect.TypeOf(QStaticText{}) // "const QStaticText &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter14drawStaticTextEiiRK11QStaticText
    // invoke: void drawStaticText(int, int, const class QStaticText &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QStaticText).qclsinst
    if false {fmt.Println(arg2)}
    C.demth_ZN8QPainter14drawStaticTextEiiRK11QStaticText(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QPainter14drawStaticTextERK7QPointFRK11QStaticText
    // invoke: void drawStaticText(const class QPointF &, const class QStaticText &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStaticText).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QPainter14drawStaticTextERK7QPointFRK11QStaticText(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN8QPainter14drawStaticTextERK6QPointRK11QStaticText
    // invoke: void drawStaticText(const class QPoint &, const class QStaticText &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStaticText).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN8QPainter14drawStaticTextERK6QPointRK11QStaticText(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawStaticText", args)
  }

}

  // proto:  void QPainter::strokePath(const QPainterPath & path, const QPen & pen);
func (this *QPainter) strokePath(args ...interface{}) () {
  // strokePath(const class QPainterPath &, const class QPen &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[0][1] = reflect.TypeOf(QPen{}) // "const QPen &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter10strokePathERK12QPainterPathRK4QPen
    // invoke: void strokePath(const class QPainterPath &, const class QPen &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPen).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QPainter10strokePathERK12QPainterPathRK4QPen(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "strokePath", args)
  }

}

  // proto:  void QPainter::drawConvexPolygon(const QPoint * points, int pointCount);
func (this *QPainter) drawConvexPolygon(args ...interface{}) () {
  // drawConvexPolygon(const class QPoint *, int)
  // drawConvexPolygon(const class QPointF *, int)
  // drawConvexPolygon(const class QPolygonF &)
  // drawConvexPolygon(const class QPolygon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter17drawConvexPolygonEPK6QPointi
    // invoke: void drawConvexPolygon(const class QPoint *, int)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN8QPainter17drawConvexPolygonEPK6QPointi(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPainter17drawConvexPolygonEPK7QPointFi
    // invoke: void drawConvexPolygon(const class QPointF *, int)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN8QPainter17drawConvexPolygonEPK7QPointFi(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN8QPainter17drawConvexPolygonERK9QPolygonF
    // invoke: void drawConvexPolygon(const class QPolygonF &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPainter17drawConvexPolygonERK9QPolygonF(this.qclsinst, arg0)
  case 3:
    // invoke: _ZN8QPainter17drawConvexPolygonERK8QPolygon
    // invoke: void drawConvexPolygon(const class QPolygon &)
    var arg0 = args[0].(QPolygon).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPainter17drawConvexPolygonERK8QPolygon(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "drawConvexPolygon", args)
  }

}

  // proto:  void QPainter::drawPath(const QPainterPath & path);
func (this *QPainter) drawPath(args ...interface{}) () {
  // drawPath(const class QPainterPath &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8drawPathERK12QPainterPath
    // invoke: void drawPath(const class QPainterPath &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QPainter8drawPathERK12QPainterPath(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "drawPath", args)
  }

}

  // proto:  QMatrix QPainter::combinedMatrix();
func (this *QPainter) combinedMatrix(args ...interface{}) () {
  // combinedMatrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter14combinedMatrixEv
    // invoke: QMatrix combinedMatrix()
    C._ZNK8QPainter14combinedMatrixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "combinedMatrix", args)
  }

}

  // proto:  void QPainter::setMatrixEnabled(bool enabled);
func (this *QPainter) setMatrixEnabled(args ...interface{}) () {
  // setMatrixEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter16setMatrixEnabledEb
    // invoke: void setMatrixEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN8QPainter16setMatrixEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setMatrixEnabled", args)
  }

}

  // proto:  void QPainter::setFont(const QFont & f);
func (this *QPainter) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QPainter7setFontERK5QFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setFont", args)
  }

}

  // proto:  void QPainter::setWindow(const QRect & window);
func (this *QPainter) setWindow(args ...interface{}) () {
  // setWindow(const class QRect &)
  // setWindow(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9setWindowERK5QRect
    // invoke: void setWindow(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QPainter9setWindowERK5QRect(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN8QPainter9setWindowEiiii
    // invoke: void setWindow(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.demth_ZN8QPainter9setWindowEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPainter", "setWindow", args)
  }

}

  // proto:  const QMatrix & QPainter::deviceMatrix();
func (this *QPainter) deviceMatrix(args ...interface{}) () {
  // deviceMatrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter12deviceMatrixEv
    // invoke: const QMatrix & deviceMatrix()
    C._ZNK8QPainter12deviceMatrixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "deviceMatrix", args)
  }

}

  // proto:  void QPainter::drawPie(int x, int y, int w, int h, int a, int alen);
func (this *QPainter) drawPie(args ...interface{}) () {
  // drawPie(int, int, int, int, int, int)
  // drawPie(const class QRect &, int, int)
  // drawPie(const class QRectF &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[0][5] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter7drawPieEiiiiii
    // invoke: void drawPie(int, int, int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    C.demth_ZN8QPainter7drawPieEiiiiii(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  case 1:
    // invoke: _ZN8QPainter7drawPieERK5QRectii
    // invoke: void drawPie(const class QRect &, int, int)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.demth_ZN8QPainter7drawPieERK5QRectii(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN8QPainter7drawPieERK6QRectFii
    // invoke: void drawPie(const class QRectF &, int, int)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN8QPainter7drawPieERK6QRectFii(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPainter", "drawPie", args)
  }

}

  // proto:  void QPainter::setWorldMatrixEnabled(bool enabled);
func (this *QPainter) setWorldMatrixEnabled(args ...interface{}) () {
  // setWorldMatrixEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter21setWorldMatrixEnabledEb
    // invoke: void setWorldMatrixEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN8QPainter21setWorldMatrixEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setWorldMatrixEnabled", args)
  }

}

  // proto:  void QPainter::QPainter(const QPainter & );
func NewQPainter(args ...interface{}) QPainter {
  return QPainter{}
}

  // proto:  void QPainter::drawPoints(const QPolygon & points);
func (this *QPainter) drawPoints(args ...interface{}) () {
  // drawPoints(const class QPolygon &)
  // drawPoints(const class QPointF *, int)
  // drawPoints(const class QPolygonF &)
  // drawPoints(const class QPoint *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPolygon{}) // "const QPolygon &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPolygonF{}) // "const QPolygonF &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QPoint{}) // "const QPoint *"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter10drawPointsERK8QPolygon
    // invoke: void drawPoints(const class QPolygon &)
    var arg0 = args[0].(QPolygon).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPainter10drawPointsERK8QPolygon(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN8QPainter10drawPointsEPK7QPointFi
    // invoke: void drawPoints(const class QPointF *, int)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN8QPainter10drawPointsEPK7QPointFi(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN8QPainter10drawPointsERK9QPolygonF
    // invoke: void drawPoints(const class QPolygonF &)
    var arg0 = args[0].(QPolygonF).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPainter10drawPointsERK9QPolygonF(this.qclsinst, arg0)
  case 3:
    // invoke: _ZN8QPainter10drawPointsEPK6QPointi
    // invoke: void drawPoints(const class QPoint *, int)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN8QPainter10drawPointsEPK6QPointi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawPoints", args)
  }

}

  // proto:  void QPainter::drawRect(int x1, int y1, int w, int h);
func (this *QPainter) drawRect(args ...interface{}) () {
  // drawRect(int, int, int, int)
  // drawRect(const class QRectF &)
  // drawRect(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8drawRectEiiii
    // invoke: void drawRect(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.demth_ZN8QPainter8drawRectEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN8QPainter8drawRectERK6QRectF
    // invoke: void drawRect(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPainter8drawRectERK6QRectF(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN8QPainter8drawRectERK5QRect
    // invoke: void drawRect(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPainter8drawRectERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "drawRect", args)
  }

}

  // proto:  QRegion QPainter::clipRegion();
func (this *QPainter) clipRegion(args ...interface{}) () {
  // clipRegion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter10clipRegionEv
    // invoke: QRegion clipRegion()
    C._ZNK8QPainter10clipRegionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "clipRegion", args)
  }

}

  // proto:  void QPainter::drawLine(const QLineF & line);
func (this *QPainter) drawLine(args ...interface{}) () {
  // drawLine(const class QLineF &)
  // drawLine(const class QPointF &, const class QPointF &)
  // drawLine(const class QLine &)
  // drawLine(int, int, int, int)
  // drawLine(const class QPoint &, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLineF{}) // "const QLineF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QLine{}) // "const QLine &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[3][3] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[4][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8drawLineERK6QLineF
    // invoke: void drawLine(const class QLineF &)
    var arg0 = args[0].(QLineF).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPainter8drawLineERK6QLineF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN8QPainter8drawLineERK7QPointFS2_
    // invoke: void drawLine(const class QPointF &, const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN8QPainter8drawLineERK7QPointFS2_(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN8QPainter8drawLineERK5QLine
    // invoke: void drawLine(const class QLine &)
    var arg0 = args[0].(QLine).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPainter8drawLineERK5QLine(this.qclsinst, arg0)
  case 3:
    // invoke: _ZN8QPainter8drawLineEiiii
    // invoke: void drawLine(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.demth_ZN8QPainter8drawLineEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  case 4:
    // invoke: _ZN8QPainter8drawLineERK6QPointS2_
    // invoke: void drawLine(const class QPoint &, const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN8QPainter8drawLineERK6QPointS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawLine", args)
  }

}

  // proto:  QPaintDevice * QPainter::device();
func (this *QPainter) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter6deviceEv
    // invoke: QPaintDevice * device()
    C._ZNK8QPainter6deviceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "device", args)
  }

}

  // proto:  void QPainter::setViewport(const QRect & viewport);
func (this *QPainter) setViewport(args ...interface{}) () {
  // setViewport(const class QRect &)
  // setViewport(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter11setViewportERK5QRect
    // invoke: void setViewport(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QPainter11setViewportERK5QRect(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN8QPainter11setViewportEiiii
    // invoke: void setViewport(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.demth_ZN8QPainter11setViewportEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPainter", "setViewport", args)
  }

}

  // proto:  void QPainter::drawTextItem(int x, int y, const QTextItem & ti);
func (this *QPainter) drawTextItem(args ...interface{}) () {
  // drawTextItem(int, int, const class QTextItem &)
  // drawTextItem(const class QPoint &, const class QTextItem &)
  // drawTextItem(const class QPointF &, const class QTextItem &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QTextItem{}) // "const QTextItem &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1][1] = reflect.TypeOf(QTextItem{}) // "const QTextItem &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[2][1] = reflect.TypeOf(QTextItem{}) // "const QTextItem &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter12drawTextItemEiiRK9QTextItem
    // invoke: void drawTextItem(int, int, const class QTextItem &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QTextItem).qclsinst
    if false {fmt.Println(arg2)}
    C.demth_ZN8QPainter12drawTextItemEiiRK9QTextItem(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QPainter12drawTextItemERK6QPointRK9QTextItem
    // invoke: void drawTextItem(const class QPoint &, const class QTextItem &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTextItem).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN8QPainter12drawTextItemERK6QPointRK9QTextItem(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN8QPainter12drawTextItemERK7QPointFRK9QTextItem
    // invoke: void drawTextItem(const class QPointF &, const class QTextItem &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTextItem).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QPainter12drawTextItemERK7QPointFRK9QTextItem(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawTextItem", args)
  }

}

  // proto:  void QPainter::save();
func (this *QPainter) save(args ...interface{}) () {
  // save()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter4saveEv
    // invoke: void save()
    C._ZN8QPainter4saveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "save", args)
  }

}

  // proto:  QTransform QPainter::combinedTransform();
func (this *QPainter) combinedTransform(args ...interface{}) () {
  // combinedTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter17combinedTransformEv
    // invoke: QTransform combinedTransform()
    C._ZNK8QPainter17combinedTransformEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "combinedTransform", args)
  }

}

  // proto:  bool QPainter::end();
func (this *QPainter) end(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter3endEv
    // invoke: bool end()
    C._ZN8QPainter3endEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "end", args)
  }

}

  // proto:  void QPainter::drawRoundRect(const QRect & r, int xround, int yround);
func (this *QPainter) drawRoundRect(args ...interface{}) () {
  // drawRoundRect(const class QRect &, int, int)
  // drawRoundRect(const class QRectF &, int, int)
  // drawRoundRect(int, int, int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = qtrt.Int32Ty(false) // "int"
  vtys[2][4] = qtrt.Int32Ty(false) // "int"
  vtys[2][5] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter13drawRoundRectERK5QRectii
    // invoke: void drawRoundRect(const class QRect &, int, int)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.demth_ZN8QPainter13drawRoundRectERK5QRectii(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN8QPainter13drawRoundRectERK6QRectFii
    // invoke: void drawRoundRect(const class QRectF &, int, int)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN8QPainter13drawRoundRectERK6QRectFii(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN8QPainter13drawRoundRectEiiiiii
    // invoke: void drawRoundRect(int, int, int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    C.demth_ZN8QPainter13drawRoundRectEiiiiii(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  default:
    qtrt.ErrorResolve("QPainter", "drawRoundRect", args)
  }

}

  // proto:  void QPainter::setWorldTransform(const QTransform & matrix, bool combine);
func (this *QPainter) setWorldTransform(args ...interface{}) () {
  // setWorldTransform(const class QTransform &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter17setWorldTransformERK10QTransformb
    // invoke: void setWorldTransform(const class QTransform &, _Bool)
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZN8QPainter17setWorldTransformERK10QTransformb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "setWorldTransform", args)
  }

}

  // proto:  void QPainter::restore();
func (this *QPainter) restore(args ...interface{}) () {
  // restore()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter7restoreEv
    // invoke: void restore()
    C._ZN8QPainter7restoreEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "restore", args)
  }

}

  // proto:  void QPainter::drawPoint(const QPoint & p);
func (this *QPainter) drawPoint(args ...interface{}) () {
  // drawPoint(const class QPoint &)
  // drawPoint(int, int)
  // drawPoint(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter9drawPointERK6QPoint
    // invoke: void drawPoint(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPainter9drawPointERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN8QPainter9drawPointEii
    // invoke: void drawPoint(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZN8QPainter9drawPointEii(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN8QPainter9drawPointERK7QPointF
    // invoke: void drawPoint(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN8QPainter9drawPointERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "drawPoint", args)
  }

}

  // proto: static QPaintDevice * QPainter::redirected(const QPaintDevice * device, QPoint * offset);
func (this *QPainter) redirected_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPainter", "redirected", args)
  }

}

  // proto:  void QPainter::shear(qreal sh, qreal sv);
func (this *QPainter) shear(args ...interface{}) () {
  // shear(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter5shearEdd
    // invoke: void shear(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN8QPainter5shearEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "shear", args)
  }

}

  // proto:  const QFont & QPainter::font();
func (this *QPainter) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter4fontEv
    // invoke: const QFont & font()
    C._ZNK8QPainter4fontEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "font", args)
  }

}

  // proto:  const QTransform & QPainter::deviceTransform();
func (this *QPainter) deviceTransform(args ...interface{}) () {
  // deviceTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter15deviceTransformEv
    // invoke: const QTransform & deviceTransform()
    C._ZNK8QPainter15deviceTransformEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "deviceTransform", args)
  }

}

  // proto:  void QPainter::resetMatrix();
func (this *QPainter) resetMatrix(args ...interface{}) () {
  // resetMatrix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter11resetMatrixEv
    // invoke: void resetMatrix()
    C._ZN8QPainter11resetMatrixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "resetMatrix", args)
  }

}

  // proto:  QPaintEngine * QPainter::paintEngine();
func (this *QPainter) paintEngine(args ...interface{}) () {
  // paintEngine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter11paintEngineEv
    // invoke: QPaintEngine * paintEngine()
    C._ZNK8QPainter11paintEngineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "paintEngine", args)
  }

}

  // proto:  bool QPainter::isActive();
func (this *QPainter) isActive(args ...interface{}) () {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter8isActiveEv
    // invoke: bool isActive()
    C._ZNK8QPainter8isActiveEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "isActive", args)
  }

}

  // proto: static void QPainter::restoreRedirected(const QPaintDevice * device);
func (this *QPainter) restoreRedirected_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPainter", "restoreRedirected", args)
  }

}

  // proto:  bool QPainter::worldMatrixEnabled();
func (this *QPainter) worldMatrixEnabled(args ...interface{}) () {
  // worldMatrixEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter18worldMatrixEnabledEv
    // invoke: bool worldMatrixEnabled()
    C._ZNK8QPainter18worldMatrixEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "worldMatrixEnabled", args)
  }

}

  // proto:  const QTransform & QPainter::transform();
func (this *QPainter) transform(args ...interface{}) () {
  // transform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter9transformEv
    // invoke: const QTransform & transform()
    C._ZNK8QPainter9transformEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "transform", args)
  }

}

  // proto: static void QPainter::setRedirected(const QPaintDevice * device, QPaintDevice * replacement, const QPoint & offset);
func (this *QPainter) setRedirected_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPainter", "setRedirected", args)
  }

}

  // proto:  QFontMetrics QPainter::fontMetrics();
func (this *QPainter) fontMetrics(args ...interface{}) () {
  // fontMetrics()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter11fontMetricsEv
    // invoke: QFontMetrics fontMetrics()
    C._ZNK8QPainter11fontMetricsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "fontMetrics", args)
  }

}

  // proto:  void QPainter::drawGlyphRun(const QPointF & position, const QGlyphRun & glyphRun);
func (this *QPainter) drawGlyphRun(args ...interface{}) () {
  // drawGlyphRun(const class QPointF &, const class QGlyphRun &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[0][1] = reflect.TypeOf(QGlyphRun{}) // "const QGlyphRun &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter12drawGlyphRunERK7QPointFRK9QGlyphRun
    // invoke: void drawGlyphRun(const class QPointF &, const class QGlyphRun &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QGlyphRun).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QPainter12drawGlyphRunERK7QPointFRK9QGlyphRun(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "drawGlyphRun", args)
  }

}

  // proto:  void QPainter::resetTransform();
func (this *QPainter) resetTransform(args ...interface{}) () {
  // resetTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter14resetTransformEv
    // invoke: void resetTransform()
    C._ZN8QPainter14resetTransformEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "resetTransform", args)
  }

}

  // proto:  const QBrush & QPainter::brush();
func (this *QPainter) brush(args ...interface{}) () {
  // brush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter5brushEv
    // invoke: const QBrush & brush()
    C._ZNK8QPainter5brushEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "brush", args)
  }

}

  // proto:  void QPainter::~QPainter();
func (this *QPainter) FreeQPainter(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPainter", "~QPainter", args)
  }

}

  // proto:  bool QPainter::begin(QPaintDevice * );
func (this *QPainter) begin(args ...interface{}) () {
  // begin(class QPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter5beginEP12QPaintDevice
    // invoke: bool begin(class QPaintDevice *)
    var arg0 = args[0].(QPaintDevice).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QPainter5beginEP12QPaintDevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "begin", args)
  }

}

  // proto:  void QPainter::scale(qreal sx, qreal sy);
func (this *QPainter) scale(args ...interface{}) () {
  // scale(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter5scaleEdd
    // invoke: void scale(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C._ZN8QPainter5scaleEdd(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "scale", args)
  }

}

  // proto:  void QPainter::setWorldMatrix(const QMatrix & matrix, bool combine);
func (this *QPainter) setWorldMatrix(args ...interface{}) () {
  // setWorldMatrix(const class QMatrix &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter14setWorldMatrixERK7QMatrixb
    // invoke: void setWorldMatrix(const class QMatrix &, _Bool)
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZN8QPainter14setWorldMatrixERK7QMatrixb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "setWorldMatrix", args)
  }

}

  // proto:  QPainterPath QPainter::clipPath();
func (this *QPainter) clipPath(args ...interface{}) () {
  // clipPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter8clipPathEv
    // invoke: QPainterPath clipPath()
    C._ZNK8QPainter8clipPathEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "clipPath", args)
  }

}

  // proto:  QPoint QPainter::brushOrigin();
func (this *QPainter) brushOrigin(args ...interface{}) () {
  // brushOrigin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter11brushOriginEv
    // invoke: QPoint brushOrigin()
    C._ZNK8QPainter11brushOriginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "brushOrigin", args)
  }

}

  // proto:  const QBrush & QPainter::background();
func (this *QPainter) background(args ...interface{}) () {
  // background()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter10backgroundEv
    // invoke: const QBrush & background()
    C._ZNK8QPainter10backgroundEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "background", args)
  }

}

  // proto:  QRect QPainter::viewport();
func (this *QPainter) viewport(args ...interface{}) () {
  // viewport()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter8viewportEv
    // invoke: QRect viewport()
    C._ZNK8QPainter8viewportEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "viewport", args)
  }

}

  // proto:  void QPainter::fillPath(const QPainterPath & path, const QBrush & brush);
func (this *QPainter) fillPath(args ...interface{}) () {
  // fillPath(const class QPainterPath &, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainterPath{}) // "const QPainterPath &"
  vtys[0][1] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8fillPathERK12QPainterPathRK6QBrush
    // invoke: void fillPath(const class QPainterPath &, const class QBrush &)
    var arg0 = args[0].(QPainterPath).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QBrush).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN8QPainter8fillPathERK12QPainterPathRK6QBrush(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "fillPath", args)
  }

}

  // proto:  bool QPainter::matrixEnabled();
func (this *QPainter) matrixEnabled(args ...interface{}) () {
  // matrixEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter13matrixEnabledEv
    // invoke: bool matrixEnabled()
    C._ZNK8QPainter13matrixEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "matrixEnabled", args)
  }

}

  // proto:  void QPainter::setTransform(const QTransform & transform, bool combine);
func (this *QPainter) setTransform(args ...interface{}) () {
  // setTransform(const class QTransform &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter12setTransformERK10QTransformb
    // invoke: void setTransform(const class QTransform &, _Bool)
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZN8QPainter12setTransformERK10QTransformb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPainter", "setTransform", args)
  }

}

  // proto:  QRect QPainter::window();
func (this *QPainter) window(args ...interface{}) () {
  // window()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter6windowEv
    // invoke: QRect window()
    C._ZNK8QPainter6windowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "window", args)
  }

}

  // proto:  void QPainter::initFrom(const QPaintDevice * device);
func (this *QPainter) initFrom(args ...interface{}) () {
  // initFrom(const class QPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "const QPaintDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter8initFromEPK12QPaintDevice
    // invoke: void initFrom(const class QPaintDevice *)
    var arg0 = args[0].(QPaintDevice).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN8QPainter8initFromEPK12QPaintDevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "initFrom", args)
  }

}

  // proto:  QFontInfo QPainter::fontInfo();
func (this *QPainter) fontInfo(args ...interface{}) () {
  // fontInfo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter8fontInfoEv
    // invoke: QFontInfo fontInfo()
    C._ZNK8QPainter8fontInfoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "fontInfo", args)
  }

}

  // proto:  void QPainter::endNativePainting();
func (this *QPainter) endNativePainting(args ...interface{}) () {
  // endNativePainting()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter17endNativePaintingEv
    // invoke: void endNativePainting()
    C._ZN8QPainter17endNativePaintingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "endNativePainting", args)
  }

}

  // proto:  void QPainter::setViewTransformEnabled(bool enable);
func (this *QPainter) setViewTransformEnabled(args ...interface{}) () {
  // setViewTransformEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter23setViewTransformEnabledEb
    // invoke: void setViewTransformEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN8QPainter23setViewTransformEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setViewTransformEnabled", args)
  }

}

  // proto:  void QPainter::setOpacity(qreal opacity);
func (this *QPainter) setOpacity(args ...interface{}) () {
  // setOpacity(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter10setOpacityEd
    // invoke: void setOpacity(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN8QPainter10setOpacityEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "setOpacity", args)
  }

}

  // proto:  const QPen & QPainter::pen();
func (this *QPainter) pen(args ...interface{}) () {
  // pen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter3penEv
    // invoke: const QPen & pen()
    C._ZNK8QPainter3penEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "pen", args)
  }

}

  // proto:  void QPainter::rotate(qreal a);
func (this *QPainter) rotate(args ...interface{}) () {
  // rotate(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter6rotateEd
    // invoke: void rotate(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN8QPainter6rotateEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPainter", "rotate", args)
  }

}

  // proto:  QRectF QPainter::clipBoundingRect();
func (this *QPainter) clipBoundingRect(args ...interface{}) () {
  // clipBoundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPainter16clipBoundingRectEv
    // invoke: QRectF clipBoundingRect()
    C._ZNK8QPainter16clipBoundingRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "clipBoundingRect", args)
  }

}

  // proto:  void QPainter::beginNativePainting();
func (this *QPainter) beginNativePainting(args ...interface{}) () {
  // beginNativePainting()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPainter19beginNativePaintingEv
    // invoke: void beginNativePainting()
    C._ZN8QPainter19beginNativePaintingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPainter", "beginNativePainting", args)
  }

}

// <= body block end
