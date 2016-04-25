import QtQuick 2.2
import QtQuick.Controls 1.1
import QtQuick.Layouts 1.0
//main.qml
ApplicationWindow {
            width: mainLayout.implicitWidth + 2 * margin
    height:mainLayout.implicitHeight + 2 * margin
    minimumWidth: mainLayout.Layout.minimumWidth + 2 * margin
    minimumHeight: mainLayout.Layout.minimumHeight + 2 * margin
    property int margin: 11 //variable to use elsewhere
    ColumnLayout {
        id:mainLayout //identity to be referred to
        anchors.fill:parent //stretch to fill parent
        anchors.margins: margin //set to variable
        GroupBox {
            id: groupBox
            title: "Fizz Buzz input"
            Layout.fillWidth: true //make sure it fills width of available space
            RowLayout {
                id: rowLayout
                anchors.fill: parent
                TextField { //this will be the input for the number to test
                    id: numberInput
                    placeholderText: fizzbuzz.number
                    Layout.fillWidth: true
                }
                Button {
                    text: "Button"
                    onClicked: fizzbuzz.number = numberInput.text
                }
            }
        }
        GroupBox {
            id: gridBox
            title: "Result"
            Layout.fillWidth: true
            TextArea {
                id: tArea
                text: fizzbuzz.status
                Layout.minimumHeight: 30
                Layout.fillHeight: true
                Layout.fillWidth: true
            }       
        }
    }
}