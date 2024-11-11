resource "aws_instance" "test" {
  ami           = "ami-0abcdef1234567890"
  instance_type = "t3.micro"

  tags = {
    Name = "HelloWorld"
  }
}

resource "aws_ec2_instance_state" "test" {
  instance_id = aws_instance.test.id
  state       = provider::assert::timecheck(["ALL"], ["January"], "America/New_York", "09:00", "17:00") ? "running" : "stopped"
}
