    JAVA虚拟机并没有规定虚拟机应该从哪里加载类，因此不同的虚拟机实现可以采用不同的实现。
    Oracle虚拟机实现为根据类路径来加载类。按照先后的顺序，类路径可以分为一下3个部分。
        1.启动类路径（bootstrap classpath） jre/lib
        2.扩展类路径（extension classpath） jre/lib/ext
        3.用户类路径（user classpath）      . -->当前目录   可以传递-classpath或-cp修改
                -classpath可以指定目录或zip或jar
                java -cp path\to\classes
                java -cp path\to\lib1.zip
                java -cp path\to\lib2.jar
                java -cp path\to\classes;lib\a.zip
                java -cp classes;lib\*
         可以通过-Xbootclasspath选项修改启动类路径
