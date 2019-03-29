TreeMap是一种带有排序功能的key-value存储结构，它是通过红黑树实现的。

public class TreeMap<K,V>
    extends AbstractMap<K,V>
    implements NavigableMap<K,V>, Cloneable, java.io.Serializable

public interface NavigableMap<K,V> extends SortedMap<K,V>

public interface SortedMap<K,V> extends Map<K,V>

SortedMap的职责是排序，而NavigableMap的职责是在排好序的集合中进行各种导航搜索的。

浏览一下NavigableMap中的部分导航方法。
// 返回小于key的第一个元素
Map.Entry<K,V> lowerEntry(K key);
... // 一系列类似方法

// 返回倒序集合
NavigableMap<K,V> descendingMap();
...

// 返回子集合，开闭区间
NavigableMap<K,V> subMap(K fromKey, boolean fromInclusive,
                             K toKey,   boolean toInclusive);
... // 一系列类似方法

TreeMap由红黑树实现，它是有序的，所以它可以反向：

TreeMap由红黑树实现，可以正序也可以逆序；
TreeMap不是线程安全的key使用：
Collections.synchronizedSortedMap(new TreeMap(...))
TreeMap中Key不能为null，Value可以为null；
TreeMap中有丰富的导航方法。
