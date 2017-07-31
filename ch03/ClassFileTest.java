package jvm.ch03;
import java.io.Serializable;

public class ClassFileTest implements Serializable{

    public static final boolean FLAG  = true;
    public static final byte BYTE = 123;
    public static final char X = 'X';
    public static final short SHORT = 12345;
    public static final int INT = 123456789;
    public static final long LONG = 12345678901L;
    public static final float PI = 3.14f;
    public static final double E = 2.71828;

    public void main(String[] args) throws RuntimeException{
        System.out.println("Hello, World!");
    }
}