// email.gp
package check

import (
	"fmt"
	"net/smtp"
	"strings"
)

//定义EMail对象
type EKMail struct {
	To       string //接收人多人以“;”分割
	From     string //发送人
	Host     string //SMTP服务器地址
	Title    string //标题
	Content  string //内容
	User     string //发送人用户名
	PassWord string //发送人密码
}

func (ekm *EKMail) SendMail() error {
	hp := strings.Split(ekm.Host, ":")
	auth := smtp.PlainAuth("", ekm.User, ekm.PassWord, hp[0])
	content_type := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + ekm.To + "\r\nFrom: " + ekm.User + "<" + ekm.User + ">\r\nSubject: " + ekm.Title + "\r\n" + content_type + "\r\n\r\n" + ekm.Content)
	send_to := strings.Split(ekm.To, ";")
	err := smtp.SendMail(ekm.Host, auth, ekm.User, send_to, msg)
	fmt.Println("Send Mail:", string(msg), "\r\n==============================")
	return err
}

func (ekm *EKMail) setTo(to string) {
	ekm.To = to
}
func (ekm *EKMail) setFrom(from string) {
	ekm.From = from
}
func (ekm *EKMail) setHost(host string) {
	ekm.Host = host
}
func (ekm *EKMail) setTitle(title string) {
	ekm.Title = title
}
func (ekm *EKMail) setContent(content string) {
	ekm.Content = content
}
func (ekm *EKMail) setUser(user string) {
	ekm.User = user
}
func (ekm *EKMail) setPassWord(passWord string) {
	ekm.PassWord = passWord
}
