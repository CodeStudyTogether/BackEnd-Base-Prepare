饿汉式:public class Singleton {

    private static Singleton instance = new Singleton();

    private Singleton() {
    }

    public static Singleton getInstance() {
        return instance;
    }
} 

不管程序是否需要这个对象的实例，总是在类加载的时候就先创建好实例，理解起来就像不管一个人想不想吃东西都把吃的先买好，如同饿怕了一样。


懒汉式:public class Singleton {

    private volatile static Singleton instance = null;

    private Singleton() {
    }

    public static synchronized Singleton getInstance() {
        if (null == instance) {
            instance = new Singleton();
        }
        return instance;
    }
}   

如果一个对象使用频率不高，占用内存还特别大，明显就不合适用饿汉式了，这时就需要一种懒加载的思想，当程序需要这个实例的时候才去创建对象，就如同一个人懒的饿到不行了才去吃东西。
