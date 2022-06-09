# Aufgaben zum Kompressionsverfahren

Im Verzeichnis/Package `dict` ist eine einfache Trie-Datenstruktur für
ein Kompressions-Wörterbuch vorgegeben.
Ihre Aufgabe besteht darin, diese zu nutzen und zu einem Kompressionsprogramm
für Textdateien weiterzuentwickeln.

## 1. Implementierung der Funktion `GetNextValue()` im Package dict

Programmieren Sie die Funktion `GetNextValue()`, die bereits vorgegeben ist.
Diese Funktion soll einen Text erwarten, der nur aus Sequenzen der Form
`"a", "ba", "bba" ...` besteht.
Sie soll diesen Text Zeichen für Zeichen konsumieren und dabei jeweils in den Trie
absteigen. Wird ein Leerzeichen gelesen oder ist der Text zu Ende,
soll die Funktion den jeweils am aktuellen Knoten gespeicherten Wert liefern und sich
beenden. Außerdem soll sie dabei den noch nicht konsumierten Teil des Textes liefern.

Jeder Aufruf von `GetNextValue()` aus einer anderen Funktion schneidet also das erste
Wort des Textes ab und liefert dessen Entsprechung aus dem Trie sowie den restlichen Text.

Schreiben Sie auch einen Test für diese Funktion.

## 2. Package für die eigentlichen Komprimierfunktionen

Legen Sie ein Package `zip` an, in dem Sie Funktionen
zum Komprimieren und Dekomprimieren eines Textes schreiben.

## 3. Dekomprimieren eines Textes

Schreiben Sie eine Funktion `Decompress()` im Package `zip`.
Diese Funktion soll einen Text erwarten, der nur aus Sequenzen der Form
`"a", "ba", "bba" ...` besteht.
Die Funktion soll den Text (mit Hilfe von `GetNextValue()`) verarbeiten und
den entsprechenden Originaltext liefern.

## 4. Zählen der Wörter in einem Text

Schreiben Sie eine Funktion `WordCounts()` im Package `zip`.
Diese Funktion soll einen String erwarten und für jedes Wort dessen im String bestimmen.
Die Funktion soll anschließend jedem Wort entsprechend seiner Häufigkeit eine Sequenz 
der Form `"a", "ba", "bba" ...`zuordnen und dann eine Map liefern, die die Wörtern auf 
ihre Sequenzen abbildet.

## 5. Aufbau des Wörterbuchs

Schreiben Sie eine Funktion `CreateDict()` im Package `dict`.
Diese Funktion soll eine Map wie aus Aufgabe 4 erwarten
und die Sequenz-Zuordnung in einem Trie speichern.
Diesen Trie soll die Funktion zurückliefern.

## 6. Schreiben des komprimierten Textes

Schreiben Sie eine Funktion `Compress()`.
Die Funktion soll einen String erwarten,
die Map und den Trie wie in den Aufgaben 4 und 5 erstellen und
dann den komprimierten Text und eine geeignete Repräsentation des Tries liefern.

**Hinweis:**
Der Trie kann sehr effizient einfach als Liste von Wörtern gespeichert werden.

## 7. Package für das Einlesen und Schreiben von Dateien

Legen Sie ein Package `files` an und schreiben Sie darin Funktionen, die eine
Textdatei öffnen und deren Inhalt als String liefert, bzw. die einen String
in eine Textdatei schreiben.

## 8. Funktionen `Pack()` und `Unpack()`

Verwenden Sie die bisherigen Funktionen, um Funktionen zu schreiben, die eine Textdatei
einlesen und eine entsprechende ge- oder entpackte Version davon schreiben.
