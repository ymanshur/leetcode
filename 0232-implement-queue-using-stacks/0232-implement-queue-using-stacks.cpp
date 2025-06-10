class MyQueue {
public:
    /*
     A simulation:
     push(1) -> 1
     push(2) -> 1 2
     push(3) -> 1 2 3
     pop()   -> 1 2 3 transform to 3 2 1; then we'll get 1 on the top
    */
    
    stack<int> myQueue;

    MyQueue() {

    }
    
    void push(int x) {
        stack<int> tempQueue;
        while (!myQueue.empty()) {
            tempQueue.push(myQueue.top());
            myQueue.pop();
        }

        myQueue.push(x);
        while (!tempQueue.empty()) {
            myQueue.push(tempQueue.top());
            tempQueue.pop();
        }
    }
    
    int pop() {
        int x = myQueue.top();
        myQueue.pop();
        return x;
    }
    
    int peek() {
        return myQueue.top();
    }
    
    bool empty() {
        return myQueue.empty();
    }
};

/**
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue* obj = new MyQueue();
 * obj->push(x);
 * int param_2 = obj->pop();
 * int param_3 = obj->peek();
 * bool param_4 = obj->empty();
 */