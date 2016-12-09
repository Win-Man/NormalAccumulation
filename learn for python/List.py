# coding=utf-8

class Node(object):
    def __init__(self,value):
        self.value = value
        self.next = None


class List(object):
    def __init__(self,head=None):
        self.head = head
        self.__size = 0

    def append(self,value):
        new_node = Node(value)
        current = self.head
        if self.head:
            while current.next:
                current = current.next
            current.next = new_node
        else:
            self.head = new_node
        self.__size  = self.__size + 1


    def insert(self,pos,value):
        if pos > self.__size:
            raise Exception('The insert pos is out of range')
        if pos == 0:
            new_node = Node(value)
            new_node.next = self.head
            self.head = new_node
        else:
            current = self.head
            for i in range(1,pos):
                current = current.next
            new_node = Node(value)
            new_node.next = current.next
            current.next = new_node
        self.__size = self.__size + 1

    # delete the first node of value
    def delete_one(self,value):
        current = self.head
        pre = self.head
        while current:
            if current.value == value:
                if current == pre:
                    self.head = self.head.next
                    current = self.head
                    pre = self.head
                else:
                    pre.next = current.next
                    pre = current
                    current = current.next
                break
                self.__size = self.__size - 1
            else:
                pre = current
                current = current.next


    #delete all nodes of value
    def delete_all(self,value):
        current = self.head
        pre = self.head
        while current:
            if current.value == value:
                if current == pre:
                    self.head = self.head.next
                    current = self.head
                    pre = self.head
                else:
                    pre.next = current.next
                    pre = current
                    current = current.next
                self.__size = self.__size - 1
            else:
                pre = current
                current = current.next


    def size(self):
        return self.__size

    def show(self):
        current = self.head
        if self.head:
            print current.value
            while current.next:
                current = current.next
                print '->{0}'.format(current.value)


if __name__ == '__main__':
    l = List()
    l.insert(0,1)
    l.insert(1,2)
    l.insert(1,3)
    l.insert(1,4)
    l.show()