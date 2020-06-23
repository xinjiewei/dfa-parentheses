import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.HashMap;

/**
 * Created by Jack on 2018/12/10.
 */
public class DFAParentheses {
    public static void main(String[] args) {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        String input = null;

        try {
            input = br.readLine();
            new MatchParenthes().match(input);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}

class MatchParenthes {
    // 因int默认是0,所以0-6状态用1-7表示,s[0]数组用于计数
    public static int[][] s = new int[8][128];

    public static HashMap<Character, Character> validChar = new HashMap();

    public static HashMap<Character, Character> leftChar = new HashMap();

    static {
        validChar.put('{', '{');
        validChar.put('}', '}');
        validChar.put('[', '[');
        validChar.put(']', ']');
        validChar.put('(', '(');
        validChar.put(')', ')');

        leftChar.put('}', '{');
        leftChar.put(']', '[');
        leftChar.put(')', '(');
    }

    public MatchParenthes() {
        // 初始化状态转化表

        // 第二列
        for (int i = 1; i < s.length; i++) {
            s[i]['{'] = 2;
        }
        // 第三列
        s[2]['}'] = 3;
        s[3]['}'] = 3;
        s[5]['}'] = 3;
        s[7]['}'] = 3;
        // 第四列
        for (int i = 1; i < s.length; i++) {
            s[i]['['] = 4;
        }
        // 第五列
        s[3][']'] = 5;
        s[4][']'] = 5;
        s[5][']'] = 5;
        s[7][']'] = 5;
        // 第六列
        for (int i = 1; i < s.length; i++) {
            s[i]['('] = 6;
        }
        // 第七列
        s[3][')'] = 7;
        s[5][')'] = 7;
        s[6][')'] = 7;
        s[7][')'] = 7;
    }

    public boolean match(String input) {
        int flag = 1; // 默认状态
        StringBuffer buf = new StringBuffer(); // 缓冲区
        for (int i = 0; i < input.length(); i++) {
            char v = input.charAt(i);

            if (!validChar.containsKey(v)) {
                // 忽略其它字符
                continue;
            }

            buf.append(v);

            int f = s[flag][v];
            if (f == 0) {
                System.out.println("no match" + buf.toString());
                return false;
            }

            flag = f;
            s[0][v]++;

            if (leftChar.containsKey(v)) {
                // 右括号不能比左括号多
                Character l = leftChar.get(v);
                if (s[0][v] > s[0][l]) {
                    System.out.println("no match" + buf.toString());
                    return false;
                }
            }
        }

        // 暂时用计数法判断是否成对
        if (s[0]['{'] == s[0]['}'] && s[0]['['] == s[0][']'] && s[0]['('] == s[0][')']) {
            System.out.println(buf.toString());
            return true;
        }
        System.out.println("not match");
        return false;
    }
}